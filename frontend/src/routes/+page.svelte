<script lang="ts">
	import { superForm } from 'sveltekit-superforms'
	import type { ActionData } from './$types.js'

	export let data
	const { form: sf, enhance, reset, capture, restore, message, errors } = superForm(data.form, {
		resetForm: false
	})

	export const snapshot = { capture, restore }
	export let form



</script>

<form method="POST" action="?/query" use:enhance>
	<button on:click={() => console.log(form?.records)}>log</button>
	{#if $message}  
		<p>{$message}</p>
	{/if}
	<label>
		<span></span>
		<input class="input w-1/2" type="text" name="town" bind:value={$sf.town}>
	</label>
	<label>
		<span></span>
		<input class="input w-1/2" type="text" name="flatType" bind:value={$sf.flatType}>
	</label>
	<label>
		<span></span>
		<input class="input w-1/2" type="text" name="price" bind:value={$sf.price}>
	</label>
	<button class="btn" type="submit">Submit</button>
</form>

