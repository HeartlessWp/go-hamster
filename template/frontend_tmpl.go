package template

const DirListTmpl = `<?xml version="1.0" encoding="iso-8859-1"?>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN" "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en">

<head>
    <title>Index of {{.Name}}</title>
    <style type="text/css">
        a,
        a:active {
            text-decoration: none;
            color: blue;
        }

        a:visited {
            color: #48468F;
        }

        a:hover,
        a:focus {
            text-decoration: underline;
            color: red;
        }

        body {
            background-color: #F5F5F5;
        }

        h2 {
            margin-bottom: 12px;
        }

        table {
            margin-left: 12px;
        }

        th,
        td {
            font: 90% monospace;
            text-align: left;
        }

        th {
            font-weight: bold;
            padding-right: 14px;
            padding-bottom: 3px;
        }

        td {
            padding-right: 14px;
        }

        td.s,
        th.s {
            text-align: right;
        }

        div.list {
            background-color: white;
            border-top: 1px solid #646464;
            border-bottom: 1px solid #646464;
            padding-top: 10px;
            padding-bottom: 14px;
        }

        div.foot {
            font: 90% monospace;
            color: #787878;
            padding-top: 4px;
        }
    </style>
</head>

<body>
    <h2>Index of {{.Name}}</h2>
    <div class="list">
        <table summary="Directory Listing" cellpadding="0" cellspacing="0">
            <thead>
                <tr>
                    <th class="n">Name</th>
                    <th class="t">Type</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td class="n"><a href="../">Parent Directory</a>/</td>
                    <td class="t">Directory</td>
                </tr>
                {{range .Children_dir}}
                <tr>
                    <td class="n"><a href="{{.}}/">{{.}}/</a></td>
                    <td class="t">Directory</td>
                </tr>
                {{end}}
                {{range .Children_files}}
                <tr>
                    <td class="n"><a href="{{.}}">{{.}}</a></td>
                    <td class="t">&nbsp;</td>
                </tr>
                {{end}}
            </tbody>
        </table>
    </div>
    <div class="foot">{{.ServerUA}}</div>
</body>

</html>`