<script>
	let name = 'Pedro'
	let age = 24
	let residence = 'Boston'
	let birthplace = 'New York'
	let occupation = 'Archeologist'
	
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
		
		const investigator = await res.json()
		result = JSON.stringify(investigator)

		if(res.ok) {
			return investigator
		} else {
			throw new Error(investigator)
		}
	}

		let promise = doPost()

		function handleClick() {
			promise = doPost() 
		}

</script>

<main>
	<input bind:value={name} />
	<input bind:value={age} />
	<input bind:value={residence} />
	<input bind:value={birthplace} />
	<input bind:value={occupation} />
	<button type="button" on:click={handleClick}>
		Create Investigator
	</button>
	<p>
		Result:
	</p>

{#await promise}
	<p>...waiting</p>
{:then investigator}
	<p>The number is {investigator.investigator.str}</p>
{:catch error}
	<p style="color: red">{error.message}</p>
{/await}
</main>

<style>

</style>
