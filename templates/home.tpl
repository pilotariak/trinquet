<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>Trinquet</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="Pilotariak - Trinquet">
    <meta name="author" content="Pilotariak">
    <link href="/static/css/readable.min.css" rel="stylesheet">
    <link href="/static/css/font-awesome.min.css" rel="stylesheet">
    <style type="text/css">
      body {
        padding-top: 70px;
      }
      h1 > span > a, h1 > span > a:hover, h1 > span > a:visited {
          color : inherit;
          font-size: 16px;
          margin-left: 10px;
      }
    </style>
    <!-- Le HTML5 shim, for IE6-8 support of HTML5 elements -->
    <!--[if lt IE 9]>
      <script src="/static/js/html5shiv.min.js"></script>
    <![endif]-->

  </head>

  <body>
    <div class="container">

      <div class="page-header">
          <h1><img src="/static/img/trinquet.png" width="50px"/>&nbsp;Trinquet </h1>
      </div>

    </div>

    <footer>
      <div class="container">
	      Trinquet <a href="/">v{{ .Version }}</a> -
	      Copyright &copy; 2019 Pilotariak
      </div>
    </footer>

  </body>
</html>
