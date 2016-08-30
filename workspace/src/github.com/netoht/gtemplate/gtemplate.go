package main

import (
    "strings"
    "os"
    "bytes"
    "io/ioutil"
    "log"
    "text/template"
    "github.com/aymerick/raymond"
)

const SEPARATOR = "="
const ARG_TEMPLATE = "g_template"
const ARG_OUTPUT = "g_output"
const ARG_TYPE = "g_type"

func main() {
    args := getArgs()
    render(args)
}

func getArgs() map[string]string {
    var args = make(map[string]string)
    for _, arg := range os.Args {
        if !strings.Contains(arg, SEPARATOR) {
            continue
        }
        pair := strings.Split(arg, SEPARATOR)
        key := pair[0]
        value := pair[1]
        if _, ok := args[key]; !ok {
            args[key] = value
        }
    }
    return args
}

func render(args map[string]string) {
    template := args[ARG_TEMPLATE]
    output := args[ARG_OUTPUT]
    ttype := args[ARG_TYPE]
    validateTemplate(template)
    validateOutput(output)
    data := renderTemplate(ttype, template, args)
    writeOutput(output, data)
}

func validateTemplate(template string) {
    if template == "" {
        log.Printf("ERROR '%v' argument is required", ARG_TEMPLATE)
        os.Exit(1)
    }
    if _, err := os.Stat(template); os.IsNotExist(err) {
        log.Printf("ERROR '%v' argument is invalid: No such file.", ARG_TEMPLATE)
        os.Exit(1)
    }
    file, err := os.OpenFile(template, os.O_RDONLY, 0666)
    if err != nil && os.IsPermission(err) {
        log.Printf("ERROR '%v' argument is invalid: Read permission denied.", ARG_TEMPLATE)
        file.Close()
        os.Exit(1)
    }
    file.Close()
}

func validateOutput(output string) {
    if output == "" {
        log.Printf("ERROR '%v' argument is required", ARG_OUTPUT)
        os.Exit(1)
    }
    file, err := os.OpenFile(output, os.O_WRONLY, 0666)
    if err != nil && os.IsPermission(err) {
        log.Printf("ERROR '%v' argument is invalid: Read permission denied.", ARG_OUTPUT)
        file.Close()
        os.Exit(1)
    }
    file.Close()
}

func renderTemplate(ttype string, template string, args map[string]string) string {
    if ttype == "" {
        ttype = "handlebars"
    }
    switch ttype {
    case "text_template":
        return renderGoTextTemplate(template, args)
    case "handlebars":
        return renderHandlebarsTemplate(template, args)
    default:
        log.Fatal("ERROR Type template. Available: handlebars, text_template")
        return ""
    }
}

func renderGoTextTemplate(templatePath string, args map[string]string) string {
    tpl, err := template.ParseFiles(templatePath)
    if err != nil {
        log.Println("ERROR parsing", err)
        os.Exit(1)
    }
    data := new(bytes.Buffer)
    err = tpl.Execute(data, args)
    if err != nil {
        log.Println("ERROR executing", err)
        os.Exit(1)
    }
    return data.String()
}

func renderHandlebarsTemplate(template string, args map[string]string) string {
    source, _ := ioutil.ReadFile(template)
    data, err := raymond.Render(string(source), args)
    if err != nil {
        log.Println("ERROR parsing template:", err)
        os.Exit(1)
    }
    return data
}

func writeOutput(output string, data string) {
    err := ioutil.WriteFile(output, []byte(data), 0666)
    if err != nil {
        log.Fatal(err)
    }
}
