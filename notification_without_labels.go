{{ define "tgshortbody" }}{{ range . }} <b>#{{ or .Labels.alertname }}</b>
{{ if gt (len .Annotations) 0 }}
<b>Summary</b>: {{.Annotations.summary}}
<b>{{.Annotations.description}}</b>{{ end }}
{{ with .PanelURL }}<a href="{{ . }}">Panel</a>    {{ end }}{{ with .DashboardURL }}<a href="{{ . }}">Dashboard</a>{{ end }}{{ with .GeneratorURL }} <a href="{{ . }}"> </a> {{ end }}{{ with .SilenceURL }} <a href="{{ . }}"></a> {{ end }}{{ end }}
{{ end }}
{{ define "tgshort" }}{{ with .Alerts.Firing }}🔥[Alerting]{{ template "tgshortbody" . }}{{ end }}{{ with .Alerts.Resolved }}✅[ОК]{{ template "tgshortbody" . }}{{ end }}{{ end }}
