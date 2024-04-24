<script lang="ts">
	import { superForm } from 'sveltekit-superforms'

	export let data
	const { form, enhance } = superForm(data.form, {
		resetForm: false,
	})

	console.log(data.records)
</script>

<form method="POST" use:enhance>
	<label>
		<span></span>
		<input class="input w-1/2" type="text" bind:value={$form.town}>
	</label>
	<label>
		<span></span>
		<input class="input w-1/2" type="text" bind:value={$form.flatType}>
	</label>
	<label>
		<span></span>
		<input class="input w-1/2" type="text" bind:value={$form.price}>
	</label>
	<button class="btn">Submit</button>
</form>
