			<div class="col-xs-5 col-xs-offset-1">
				<p>&nbsp;</p>
				<div>
					<p>
						<input id="toggleServer" type="checkbox" {{if .dns_server}}checked{{end}} data-size="large" data-toggle="toggle" data-on="" data-off="" data-onstyle="success" data-offstyle="danger" onChange="dnsServerChange(this)">
						<span style="font-size: 18px">DNS Server</span>
					</p>
				</div>
				<div id="divFilter">
					<p>
						<input id="toggleFilter" type="checkbox" {{if .dns_filter}}checked{{end}} data-size="large" data-toggle="toggle" data-on="" data-off="" data-onstyle="success" data-offstyle="danger" onChange="dnsFilterChange(this)">
						<span style="font-size: 18px">DNS Filter</span>
					</p>
				</div>
			</div>
			<div class="col-xs-6">
				<p class="text-center"><a href="/hax?kexploit10"><button type="button" class="btn btn-danger btn-lg btn-action">Kernel exploit</button></a></p>
				<p>&nbsp;</p>
				<p class="text-center"><a href="/hax?launcher"><button type="button" class="btn btn-primary btn-lg btn-action">Homebrew launcher</button></a></p>
				<p class="text-center"><a href="/hax?loadiine"><button type="button" class="btn btn-primary btn-lg btn-action">Loadiine</button></a></p>
			</div>
			<script>
				function dnsServerChange(checkbox) {
					if (checkbox.checked) {
						$('#divFilter').show();
					} else {
						$('#divFilter').hide();
					}
					$.ajax({
						url: "/api/dns?dns_server=" + checkbox.checked,
						method: "post",
					});
				}
				function dnsFilterChange(checkbox) {
					$.ajax({
						url: "/api/dns?dns_filter=" + checkbox.checked,
						method: "post",
					});
				}
				$(function() {
					if ($('#toggleServer').prop('checked')) {
						$('#divFilter').show();
					} else {
						$('#divFilter').hide();
					}
				})
			</script>