package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var (
	data    ipHolder
	quit    bool
	gotData bool
	uiStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		PaddingRight(10).PaddingLeft(10).
		Align(lipgloss.Center).Faint(false)
)

func newModel() model {
	return model{
		spinner: spinner.Model{Spinner: randomSpinner()},
	}
}

type model struct {
	spinner  spinner.Model
	madeCall bool
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	if !m.madeCall {
		go m.getData(&data, &gotData)
		m.madeCall = true
	}

	if quit {
		cmd = tea.Quit
	} else {
		m.spinner, cmd = m.spinner.Update(msg)
	}

	return m, cmd
}

func (m model) View() string {

	var ui string

	if gotData {
		ui += dataString(data)
		quit = true
	} else {
		ui += m.spinner.View() + " fetching"
	}

	return uiStyle.Render(ui) + "\n"

}

func (m model) getData(h *ipHolder, q *bool) {
	*h = getIp()
	*q = true
}

func randomSpinner() spinner.Spinner {
	spinners := [...]spinner.Spinner{
		//spinner.Dot,
		//spinner.Jump,
		//spinner.Line,
		//spinner.Globe,
		//spinner.Moon,
		spinner.Meter,
		//spinner.Pulse,
		//spinner.Monkey,
		//spinner.Points,
		//spinner.Hamburger,
		//spinner.MiniDot,
		//spinner.Ellipsis,
        gamePadSpinner(),
        circleSliceSpinner(),
        pleaseWaitSpinner(),
	}

	return spinners[rand.Intn(len(spinners))]

}

func dataString(data ipHolder) (ui string) {
	ui += fmt.Sprintf("IpAddress(ipv%d): %s\n", data.IpVersion, data.IpAddress)
	ui += fmt.Sprintf("Proxy: %t\n", data.IsProxy)
	ui += fmt.Sprintf("Latitude and Longitude: %f, %f\n", data.Latitude, data.Longitude)
	ui += fmt.Sprintf("CityName: %s\n", data.CityName)
	ui += fmt.Sprintf("RegionName: %s\n", data.RegionName)
	ui += fmt.Sprintf("ZipCode: %s\n", data.ZipCode)
	ui += fmt.Sprintf("Country: %s(%s)\n", data.Country, data.CountryCode)
	ui += fmt.Sprintf("Continent: %s(%s)\n", data.Continent, data.ContinentCode)
	ui += fmt.Sprintf("TimeZone: %s", data.TimeZone)
	return
}

func termSize() [2]int {
	x, y, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}
	return [2]int{x, y}
}

func TermHeight() int {
	return termSize()[1]
}

func TermWidth() int {
	return termSize()[0]
}

func gamePadSpinner() spinner.Spinner {
    return spinner.Spinner{
        Frames: []string{"󰸴", "󰸵", "󰸸", "󰸷"},
        FPS: time.Second / 4,
    }
}

func circleSliceSpinner() spinner.Spinner {
    return spinner.Spinner{
        Frames: []string{"󰪞", "󰪟", "󰪠", "󰪡", "󰪢", "󰪣", "󰪤", "󰪥"},
        FPS: time.Second / 8,
    }
}

func pleaseWaitSpinner() spinner.Spinner {
    return spinner.Spinner{
        Frames: []string{
            "please wait",
            "Please wait",
            "pLease wait",
            "plEase wait",
            "pleAse wait",
            "pleaSe wait",
            "pleasE wait",
            "please Wait",
            "please wAit",
            "please waIt",
            "please waiT",
        },
        FPS: time.Second / 11,
    }
}
