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

		export function handleClick() {
			promise = doPost() 
		}
</script>

<form>
  <div class="mb-3">
    <label for="name" class="form-label">Name</label>
	<input bind:value="{name}" type="text" class="form-control" id="name" aria-describedby="name">
  </div>
 <div class="mb-3">
    <label for="age" class="form-label">Age</label>
	<input bind:value="{age}" type="number" class="form-control" id="age" aria-describedby="age">
  </div>
   <div class="mb-3">
    <label for="residence" class="form-label">Residence</label>
	<input bind:value="{residence}" type="text" class="form-control" id="residence" aria-describedby="residence">
  </div>
   <div class="mb-3">
    <label for="birthplace" class="form-label">Birthplace</label>
	<input bind:value="{birthplace}" type="text" class="form-control" id="birthplace" aria-describedby="birthplace">
  </div>
 <div class="mb-3">
    <label for="occupation" class="form-label">Occupation</label>
	<input bind:value="{occupation}" type="text" class="form-control" id="occupation" aria-describedby="occupation">
  </div>
  <button on:click="{handleClick}" type="button" class="btn btn-primary" data-toggle="collapse">Create</button>
</form>

{#await promise}
	<p>...waiting</p>
{:then investigator}
	<div>
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

		{investigator.investigator.description.str_description}
		{investigator.investigator.description.app_description}
		{investigator.investigator.description.con_description}
		{investigator.investigator.description.int_description}
		{investigator.investigator.description.siz_description}
		{investigator.investigator.description.pow_description}
		{investigator.investigator.description.edu_description}
		{investigator.investigator.description.dex_description}
	</div>
{:catch error}
	<p style="color: red">{error.message}</p>
{/await}

