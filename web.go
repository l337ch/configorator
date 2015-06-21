package main

import (
  "net"
  "net/http"
  "text/template"
  "path/filepath"
)

var baseConfig *template.Template

type CloudConfig struct {
  Config  string
  Content interface{}
}

type HostConfig struct {
  HostIpAddr string
}

//type ipAddress string

func renderConfig(w http.ResponseWriter, r *http.Request) {
  ipAddr, _, _ := net.SplitHostPort(r.RemoteAddr)
  var err error = baseConfig.Execute(w, HostConfig{
    ipAddr,
  })
  if err != nil {
    panic(err)
  }
}

func webListener(templateDir string) {
  templatePath := templateDir + "/base-config.tmpl"
  baseConfigTemplateFile, _ := filepath.Abs(templatePath) 
  baseConfig = template.Must(template.ParseFiles(baseConfigTemplateFile))
  http.HandleFunc("/cloud-configs/", renderConfig)
  http.ListenAndServe(":80", nil)
}