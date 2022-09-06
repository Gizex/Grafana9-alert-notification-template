{{ define "tgshortbody" }}{{ range . }} <b>{{ or .Labels.alertname }}</b>
{{ range .Annotations.SortedPairs }}<b>{{ .Name }}:</b> {{ .Value }}
{{ end }}
{{ with .ValueString }}{{ reReplaceAll "[[][^]]*metric='{?([^}']*)}?'[^]]*value=([0-9]*([.][0-9]{,3})?)[^]]*](, )?" "<b>$1:</b> <code> $2 </code> \n" . }}
{{ end }}{{ with .PanelURL }}<a href="{{ . }}">Panel</a>    {{ end }}{{ with .DashboardURL }}<a href="{{ . }}">Dashboard</a>{{ end }}{{ with .GeneratorURL }} <a href="{{ . }}"> </a> {{ end }}{{ with .SilenceURL }} <a href="{{ . }}"></a> {{ end }}
{{ end }}{{ end }}
{{ define "tgshort" }}{{ with .Alerts.Firing }}🔥[Alerting]{{ template "tgshortbody" . }}{{ end }}{{ with .Alerts.Resolved }}✅[ОК]{{ template "tgshortbody" . }}{{ end }}{{ end }}