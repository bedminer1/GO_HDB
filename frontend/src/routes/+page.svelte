<script lang="ts">
	import { superForm } from 'sveltekit-superforms'
	import { Paginator, ProgressBar } from '@skeletonlabs/skeleton'
    import type { PaginationSettings } from '@skeletonlabs/skeleton'

	export let data
	const { form: sf, enhance, reset, capture, restore, message, errors } = superForm(data.form, {
		resetForm: false
	})

	export const snapshot = { capture, restore }
	export let form
	let records: HDBRecord[] = []
	$: if (form) records = form.records!

	// PAGINATION
    let paginationSettings = {
        page: 0,
        limit: 10,
        size: records.length,
        amounts: [3, 5, 10, 15]
    } satisfies PaginationSettings

    $: paginationSettings.size = records.length

    $: paginatedSource = records.slice(
        paginationSettings.page * paginationSettings.limit,
        paginationSettings.page * paginationSettings.limit + paginationSettings.limit
    );


</script>

<form method="POST" action="?/query" use:enhance>
	<label>
		<span></span>
		<input class="input w-1/2" type="text" name="town" bind:value={$sf.town}>
		{#if $errors.town}
			<span>{$errors.town}</span>
		{/if}
	</label>
	<label>
		<span></span>
		<input class="input w-1/2" type="text" name="flatType" bind:value={$sf.flatType}>
		{#if $errors.flatType}
			<span>{$errors.flatType}</span>
		{/if}
	</label>
	<label>
		<span></span>
		<input class="input w-1/2" type="text" name="price" bind:value={$sf.price}>
		{#if $errors.price}
			<span>{$errors.price}</span>
		{/if}
	</label>
	<button class="btn" type="submit">Submit</button>
</form>

{#if $message}<p>{$message}</p>{/if}

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
			{#each paginatedSource as record}
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

	<Paginator 
            bind:settings={paginationSettings}
        />
</div>
{/if}

