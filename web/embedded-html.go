package web

const mainPageHTML = `<html>

	<textarea id="code" cols=80 rows=4></textarea>
	<button id="execute">Execute</button>
	<div id="progress"></div>
	<div id="output">
	Result will show up here.
	</div>
	<div id="error">
	Errors will show up here.
	</div>

	<style>
		#code, #progress, #output, #error {
			font-family: monospace;
			font-size: 11pt;
		}
		#error {
			color: red;
		}
	</style>

	<script>
	  // TODO(xiaq): Stream results.
		var $code = document.getElementById('code'),
		    $execute = document.getElementById('execute'),
				$progress = document.getElementById('progress'),
				$output = document.getElementById('output'),
				$error = document.getElementById('error');

    $execute.addEventListener('click', function() {
			var code = $code.value;
			var req = new XMLHttpRequest();

			$progress.innerText = 'executing...';
			req.onloadend = function() {
				$progress.innerText = 'executed ' + code;
			};
			req.onload = function() {
				$output.innerText = req.responseText;
				$error.innerText = '';
			};
			req.onerror = function() {
				$output.innerText = '';
				$error.innerText = req.responseText || 'unknown error';
			};
			req.open('POST', '/execute');
			req.send(code);
		});
	</script>

</html>
`