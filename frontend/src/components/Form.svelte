<script>
	import Hidden from './Hidden.svelte'
	let name = ''
	let age = null
	let residence = ''
	let birthplace = ''
	let occupation = ''
	
	export async function doPost () {
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
	  <input bind:value={occupation} class="input" type="text" placeholder="Investigator's occupation ie: Archeologist">
  </div>
</div> 

 <div class="control">
	 <button on:click={handleClick} class="button is-link">Generate</button>
 	<button on:click={child.show} class="button">Show</button>
 </div>



 <section class="section">

 </section>

{#await promise}
	<p>...waiting</p>
{:then investigator}
 
<Hidden bind:this={child} on:show={e => child.shown = e.detail}>

<div class="table-container">
  <table class="table is-striped is-fullwidth">
	<tbody>
		<tr>Name: {investigator.investigator.name}</tr>  
		<tr>Age: {investigator.investigator.age}</tr>  
		<tr>Age: {investigator.investigator.age}</tr>
		<tr>Residence: {investigator.investigator.residence}</tr>
		<tr>Birthplace: {investigator.investigator.birthplace}</tr>
		<tr>Occupation: {investigator.investigator.occupation}</tr>
		<tr>Strenght: {investigator.investigator.str}</tr>
		<tr>Constitution: {investigator.investigator.con}</tr>
		<tr>Power: {investigator.investigator.pow}</tr>
		<tr>Dexterity: {investigator.investigator.dex}</tr>
		<tr>Appearance: {investigator.investigator.app}</tr>
		<tr>Size: {investigator.investigator.siz}</tr>
		<tr>Intelligence: {investigator.investigator.int}</tr>
		<tr>Education: {investigator.investigator.edu}</tr>
		<tr>Luck: {investigator.investigator.luck}</tr>
		<tr>MP: {investigator.investigator.mp}</tr>
		<tr>Damage Bonus: {investigator.investigator.db}</tr>
		<tr>Build: {investigator.investigator.build}</tr>
		<tr>HP: {investigator.investigator.hp}</tr>
		<tr>Sanity: {investigator.investigator.san}</tr>
		<tr>Move Rate: {investigator.investigator.mv}</tr>
		</tbody>
 </table>
</div>



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

	<div>
		{investigator.investigator.description.str_description}
		{investigator.investigator.description.app_description}
		{investigator.investigator.description.con_description}
		{investigator.investigator.description.int_description}
		{investigator.investigator.description.siz_description}
		{investigator.investigator.description.pow_description}
		{investigator.investigator.description.edu_description}
		{investigator.investigator.description.dex_description}
	</div>
</Hidden>
{/await}

