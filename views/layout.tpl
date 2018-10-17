<!DOCTYPE html>
<html lang="en">
<head>
	<title>goHaxServer</title>

	<meta charset="utf-8">
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">

	<link rel="icon" href="/static/img/favicon.ico" />

	<link type="text/css" rel="stylesheet" href="/static/css/bootstrap.min.css" />
	<link type="text/css" rel="stylesheet" href="/static/css/gohaxserver.css" />
	<link type="text/css" rel="stylesheet" href="/static/css/bootstrap-toggle.min.css" rel="stylesheet" />
</head>

<body>
	<script src="/static/js/jquery-3.1.1.min.js"></script>
	<script src="/static/js/bootstrap.min.js"></script>
	<script src="/static/js/bootstrap-toggle.min.js"></script>

	<div class="container-fluid">
		<div class="row">
			<div>
				<div id="logo" class="col-xs-6"></div>
			</div>
		</div>
		<div class="row">&nbsp;</div>
		<div class="row">
{{.LayoutContent}}
		</div>
	</div>

	<div id="footer">goHaxServer v{{.version}}</div>
</body>
</html>
