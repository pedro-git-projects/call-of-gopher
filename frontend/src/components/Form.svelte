<script>
	let name = ''
	let age = null
	let residence = ''
	let birthplace = ''
	let occupation = ''
	
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

<div class="field">
  <label class="label">Name</label>
  <div class="control">
	  <input bind:value={name} class="input" type="text" placeholder="Investigator's name ie: Dexter Ward">
  </div>
</div> 

<div class="field">
  <label class="label">Age</label>
  <div class="control">
    <input bind:value={age} class="input" type="number" placeholder="Investigator's age ie: 31">
  </div>
</div> 

<div class="field">
  <label class="label">Residence</label>
  <div class="control">
	  <input bind:value={residence} class="input" type="text" placeholder="Investigator's residence ie: New York">
  </div>
</div> 
<div class="field">
  <label class="label">Birthplace</label>
  <div class="control">
	  <input bind:value={birthplace} class="input" type="text" placeholder="Investigator's birthplace ie: Newport">
  </div>
</div> 
<div class="field">

  <label class="label">Occupation</label>
  <div class="control">
	  <input on:click={handleClick} bind:value={occupation} class="input" type="text" placeholder="Investigator's occupation ie: Archeologist">
  </div>
</div> 

 <div class="control">
	 <button on:click={handleClick} class="button is-link">Generate</button>
 </div>


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

