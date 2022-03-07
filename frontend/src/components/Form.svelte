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
				age, residence,
				birthplace,
				occupation
			})
		})
		
		const investigator = await res.json()

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
	<p>
		Name: {investigator.investigator.name}
		Age: {investigator.investigator.age}
		Residence: {investigator.investigator.residence}
		Birthplace: {investigator.investigator.birthplace}
		Occupation: {investigator.investigator.occupation}
		Strenght: {investigator.investigator.str}
		Constitution: {investigator.investigator.con}
		Power: {investigator.investigator.pow}
		Dexterity: {investigator.investigator.dex}
		Appearance: {investigator.investigator.app}
		Size: {investigator.investigator.siz}
		Intelligence: {investigator.investigator.int}
		Education: {investigator.investigator.edu}
		Luck: {investigator.investigator.luck}
		MP: {investigator.investigator.mp}
		Damage Bonus: {investigator.investigator.db}
		Build: {investigator.investigator.build}
		HP: {investigator.investigator.hp}
		Sanity: {investigator.investigator.san}
		Move Rate: {investigator.investigator.mv}
	</p>
	<p>
		{investigator.investigator.description.str_description}
		{investigator.investigator.description.app_description}
		{investigator.investigator.description.con_description}
		{investigator.investigator.description.int_description}
		{investigator.investigator.description.siz_description}
		{investigator.investigator.description.pow_description}
		{investigator.investigator.description.edu_description}
		{investigator.investigator.description.dex_description}
	</p>
{:catch error}
	<p style="color: red">{error.message}</p>
{/await}

