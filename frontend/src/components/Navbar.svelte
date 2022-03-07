<script>
	let activeClass ="" 
	function isActive() {
		if(activeClass == "") {
			activeClass = "is-active"
		} else if(activeClass == "is-active") {
			activeClass = ""
		}
	} 
</script>

<nav class="navbar is-dark" role="navigation" aria-label="main navigation">
	<div class="navbar-brand">
		<a class="navbar-item" href="#">
			<img src="https://go.dev/images/go-logo-white.svg" width="112" height="28">
		</a>

		<a on:click={isActive} role="button"  class="navbar-burger {activeClass}" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample"> <span aria-hidden="true"></span> <span aria-hidden="true"></span> <span aria-hidden="true"></span> </a>
	</div>

	<div id="navbarBasicExample" class="navbar-menu {activeClass}">
		<div class="navbar-start">
			<a href="https://github.com/pedro-git-projects/go-of-cthullu" target="_blank" class="navbar-item">
				Source Code
			</a>

			<a class="navbar-item">
				Documentation
			</a>
		</div>
		<div class="navbar-end">
		</div>
	</div>
</nav>


