<script lang="ts">
	import { superForm } from 'sveltekit-superforms'
	import type { ActionData } from './$types.js'

	export let data
	const { form: sf, enhance, reset, capture, restore, message, errors } = superForm(data.form, {
		resetForm: false
	})

	export const snapshot = { capture, restore }
	export let form
	let records: HDBRecord[] = []
	$: if (form) records = form.records!


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

{#if form?.records}
<div>
	<table class="table w-2/3">
		<thead>
			<th>Month</th>
			<th>Town</th>
			<th>Flat Type</th>
			<th>Lease Start Year</th>
			<th>Remaining Lease</th>
			<th>Resale Price</th>
		</thead>
		<tbody>
			{#each form.records as record}
				<tr>
					<td>{record.month}</td>
					<td>{record.town}</td>
					<td>{record.flatType}</td>
					<td>{record.leaseStart}</td>
					<td>{record.remainingLease}</td>
					<td>{record.price}</td>
				</tr>

			{/each}
		</tbody>
	</table>
</div>
{/if}

