package main

import (
    "gpio"
    "flag"
    "log"
)

var gpio_num string
var on_off string

func init() {
    flag.StringVar(&gpio_num, "gpio", "17", "gpio number")
    flag.StringVar(&on_off, "on-off", "on", "switch on off")
}

func main() {
    current_gpio := &gpio.GoGPIO{PinNumber: gpio_num}
    if on_off == "on" {
        current_gpio.Export()
        current_gpio.On()
    } else if on_off == "off" {
        current_gpio.Off()
        current_gpio.UnExport()
    } else {
        log.Fatal("not support control command")
    }
}