/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */

package winipcfg

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"net"
	"os/exec"
	"strings"
	"syscall"

	"golang.org/x/sys/windows"
)

// I wish we didn't have to do this. netiohlp.dll (what's used by netsh.exe) has some nice tricks with writing directly
// to the registry and the nsi kernel object, but it's not clear copying those makes for a stable interface. WMI doesn't
// work with v6. CMI isn't in Windows 7.
func runNetsh(cmds []string) error {
	system32, err := windows.GetSystemDirectory()
	if err != nil {
		return err
	}
	cmd := exec.Command(system32 + "\\netsh.exe") // I wish we could append (, "-f", "CONIN$") but Go sets up the process context wrong.
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return errors.New(fmt.Sprintf("runNetsh stdin pipe - %v", err))
	}
	go func() {
		defer stdin.Close()
		encoder := simplifiedchinese.GB18030.NewEncoder()
		transformedInput := transform.NewWriter(stdin, encoder)
		_, writeErr := transformedInput.Write([]byte(strings.Join(append(cmds, "exit\r\n"), "\r\n")))
		if writeErr != nil {
			fmt.Println(writeErr)
		}
	}()
	output, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(fmt.Sprintf("runNetsh run - %v", err))
	}
	// Horrible kludges, sorry.
	cleaned := bytes.ReplaceAll(output, []byte("netsh>"), []byte{})
	cleaned = bytes.ReplaceAll(cleaned, []byte("There are no Domain Name Servers (DNS) configured on this computer."), []byte{})
	cleaned = bytes.TrimSpace(cleaned)
	if len(cleaned) != 0 {
		return errors.New(fmt.Sprintf("runNetsh returned error strings.\ninput:\n%s\noutput\n:%s",
			strings.Join(cmds, "\n"), bytes.ReplaceAll(output, []byte{'\r', '\n'}, []byte{'\n'})))
	}
	return nil
}

func runNetshResult(cmds []string) (string, error) {
	system32, err := windows.GetSystemDirectory()
	if err != nil {
		return "", fmt.Errorf("runNetshResult GetSystemDirectory - %w", err)
	}
	cmd := exec.Command(system32 + "\\netsh.exe") // I wish we could append (, "-f", "CONIN$") but Go sets up the process context wrong.
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", fmt.Errorf("runNetshResult stdin pipe - %w", err)
	}
	go func() {
		defer stdin.Close()
		// 如果有中文，需要设置编码
		// 设置编码为GB18030
		encoder := simplifiedchinese.GB18030.NewEncoder()
		transformedInput := transform.NewWriter(stdin, encoder)
		_, writeErr := transformedInput.Write([]byte(strings.Join(append(cmds, "exit\r\n"), "\r\n")))
		if writeErr != nil {
			fmt.Println(writeErr)
		}
		//io.WriteString(stdin, strings.Join(append(cmds, "exit\r\n"), "\r\n"))
	}()
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(ConvertByte2String(output, GB18030))
		return "", fmt.Errorf("runNetshResult run - %w", err)
	}
	// Horrible kludges, sorry.
	cleaned := bytes.ReplaceAll(output, []byte("netsh>"), []byte{})

	return ConvertByte2String(cleaned, GB18030), nil
}

func flushDnsCmds(ifc *Interface) []string {
	return []string{
		fmt.Sprintf("interface ipv4 set dnsservers name=%d source=static address=none validate=no register=both", ifc.Index),
		fmt.Sprintf("interface ipv6 set dnsservers name=%d source=static address=none validate=no register=both", ifc.Ipv6IfIndex),
	}
}

func addDnsCmds(ifc *Interface, dnses []net.IP) []string {
	cmds := make([]string, len(dnses))
	j := 0
	for i := 0; i < len(dnses); i++ {
		if v4 := dnses[i].To4(); v4 != nil {
			cmds[j] = fmt.Sprintf("interface ipv4 add dnsservers name=%d address=%s validate=no", ifc.Index, v4.String())
		} else if v6 := dnses[i].To16(); v6 != nil {
			cmds[j] = fmt.Sprintf("interface ipv6 add dnsservers name=%d address=%s validate=no", ifc.Ipv6IfIndex, v6.String())
		} else {
			continue
		}
		j++
	}
	return cmds[:j]
}

func RunNetsh(cmds []string) (string, error) {
	return runNetshResult(cmds)
}

const (
	netshCmdTemplateEnablingInterface  = "interface set interface %s admin=enable"
	netshCmdTemplateDisablingInterface = "interface set interface %s admin=disable"
	netshCmdTemplateStatusInterface    = "interface show interface %s"
)

// 开启网卡
func EnablingInterface(interfaceName string) error {
	//netsh interface set interface "StarVPN" admin=enable
	_, err := runNetshResult([]string{fmt.Sprintf(netshCmdTemplateEnablingInterface, interfaceName)})
	return err
}

// 禁用网卡
func DisablingInterface(interfaceName string) error {
	//netsh interface set interface "StarVPN" admin=disable
	_, err := runNetshResult([]string{fmt.Sprintf(netshCmdTemplateDisablingInterface, interfaceName)})
	return err
}

type InterfaceStatus uint32

var (
	// 禁用
	INTERFACE_STATUS_DISABLED InterfaceStatus = 0
	// 启用
	INTERFACE_STATUS_ENABLED InterfaceStatus = 1
	// 已连接
	INTERFACE_STATUS_CONNECTED InterfaceStatus = 2
	// 未知
	INTERFACE_STATUS_UNKNOWN InterfaceStatus = 3
)

// 查看网卡状态 0-
func FindInterfaceStatus(interfaceName string) (InterfaceStatus, error) {
	//netsh interface show interface "StarVPN"
	result, err := runNetshResult([]string{fmt.Sprintf(netshCmdTemplateStatusInterface, interfaceName)})
	if err != nil {
		return INTERFACE_STATUS_UNKNOWN, err
	}
	lines := strings.Split(result, "\n")
	for _, line := range lines {
		// 去除空格
		line = strings.TrimSpace(line)
		// 按:分割
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}
		if strings.TrimSpace(parts[0]) == "管理状态" {
			statusStr := strings.TrimSpace(parts[1])
			if statusStr == "已禁用" {
				return INTERFACE_STATUS_DISABLED, nil
			}
		} else if strings.TrimSpace(parts[0]) == "连接状态" {
			statusStr := strings.TrimSpace(parts[1])
			if statusStr == "已连接" {
				return INTERFACE_STATUS_CONNECTED, nil
			}
		}
	}
	return INTERFACE_STATUS_ENABLED, nil
}

// 修改网卡名称
func RenameInterface(oldName, newName string) error {
	_, err := runNetshResult([]string{fmt.Sprintf("interface set interface name=\"%s\" newname=\"%s\"", oldName, newName)})
	if err != nil {
		return err
	}
	return nil
}
