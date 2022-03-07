<script>
	let name = 'Pedro'
	let age = 24
	let residence = 'Boston'
	let birthplace = 'New York'
	let occupation = 'Archeologist'
	let result = null
	
	async function doPost () {
		const res = await fetch('http://localhost:4000/v1/investigator', {
			method: 'POST',
			body: JSON.stringify({
				name,
				age,
				residence,
				birthplace,
				occupation
			})
		})
		
		const json = await res.json()
		result = JSON.stringify(json)
		investigator = json	
	}
</script>

<main>
	<input bind:value={name} />
	<input bind:value={age} />
	<input bind:value={residence} />
	<input bind:value={birthplace} />
	<input bind:value={occupation} />
	<button type="button" on:click={doPost}>
	Post it.
	</button>
	<p>
		Result:
	</p>
	<pre>
		{result}
	</pre>
	<p>
	</p>
</main>

<style>

</style>
