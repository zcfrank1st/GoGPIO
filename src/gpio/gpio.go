package main

import (
    "os"
    "log"
    "io/ioutil"
    "time"
    "fmt"
)

const (
    GPIO_VALUE_TEMPLETE = "/sys/class/gpio/gpio%s/value"
)

type GoGPIO struct {
    PinNumber string
}

func (g *GoGPIO) Export() {
    g.writeGPIOFile("/sys/class/gpio/export", g.PinNumber)
    time.Sleep(1 * time.Second)
    g.writeGPIOFile(fmt.Sprintf("/sys/class/gpio/gpio%s/direction", g.PinNumber), "out")
}

func (g *GoGPIO) On() {
    g.writeGPIOFile(fmt.Sprintf(GPIO_VALUE_TEMPLETE, g.PinNumber), "1")
}

func (g *GoGPIO) Off() {
    g.writeGPIOFile(fmt.Sprintf(GPIO_VALUE_TEMPLETE, g.PinNumber), "0")
}

func (g *GoGPIO) Status() string {
    if value, err := g.readGPIOFile(fmt.Sprintf(GPIO_VALUE_TEMPLETE, g.PinNumber)); err == nil {
        return value
    } else {
        return ""
    }
}

func (g *GoGPIO) UnExport() {
    g.writeGPIOFile("/sys/class/gpio/unexport", g.PinNumber)
}


func (g *GoGPIO) readGPIOFile(path string) (string, error){
    if value, err := ioutil.ReadFile(path); err == nil {
        return string(value), nil
    } else {
        log.Println("read gpio fatal", err)
        return "", err
    }
}

func (g *GoGPIO) writeGPIOFile(path string, value string) {
    if f, err := os.OpenFile(path, os.O_APPEND | os.O_WRONLY, os.FileMode(0777)); err == nil {
        defer f.Close()

        if _,err := f.WriteString(value); err != nil {
            log.Fatal("write gpio fatal", err)
        }
    } else {
        log.Fatal("manipulate gpio, fatal : ", err)
    }
}