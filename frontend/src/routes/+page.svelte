<script lang="ts">
	import { superForm } from 'sveltekit-superforms'
	import { Paginator } from '@skeletonlabs/skeleton'
    import type { PaginationSettings } from '@skeletonlabs/skeleton'
	import BarChart from '$lib/components/BarChart.svelte'

	// SUPERFORM INIT
	export let data
	const { form: sf, enhance, capture, restore, message, errors } = superForm(data.form, {
		resetForm: false
	})
	export const snapshot = { capture, restore }

	// DATA FROM QUERY
	export let form
	let records: HDBRecord[] = []
	let years: string[] = []
	let meanData: number[] = []
	let countData: number[] = []
	$: {
		if (form) {
			records = form.records!
			years = form.years!
			meanData = form.meanData!
			countData = form.countData!
		}
	} 

	// TABLE/PAGINATION
    let paginationSettings = {
        page: 0,
        limit: 10,
        size: records?.length,
        amounts: [3, 5, 10, 15]
    } satisfies PaginationSettings

    $: paginationSettings.size = records?.length

	let paginatedSource: HDBRecord[]
    $: if (records) paginatedSource = records.slice(
        paginationSettings.page * paginationSettings.limit,
        paginationSettings.page * paginationSettings.limit + paginationSettings.limit
    );

	// CHART
	$: meanDataSet = {
		labels: years,
		datasets: [{
			label: "Annual Mean",
			data: meanData,
			borderWidth: 1,
            backgroundColor: '#DCC7EA',
		}],
	}
	$: countDataSet = {
		labels: years,
		datasets: [{
			label: "Number of Listings",
			data: countData,
			borderWidth: 1,
            backgroundColor: '#DCC7EA',
		}],
	}
</script>

<div class="w-full min-h-screen flex justify-center flex-col gap-4 items-center">
	<form method="POST" action="?/query" use:enhance class="w-1/2 text-center flex justify-center gap-2 mb-8">
		<label>
			<span>Town</span>
			<input class="input" type="text" name="town" bind:value={$sf.town} placeholder="Town">
			{#if $errors.town}
				<span>{$errors.town}</span>
			{/if}
		</label>
		<label>
			<span>Flat Type</span>
			<input class="input" type="text" name="flatType" bind:value={$sf.flatType} placeholder="Flat Type">
			{#if $errors.flatType}
				<span>{$errors.flatType}</span>
			{/if}
		</label>
		<label>
			<span>Max Price (K)</span>
			<input class="input" type="text" name="price" bind:value={$sf.price} placeholder="Max Price">
			{#if $errors.price}
				<span>{$errors.price}</span>
			{/if}
		</label>
		<button class="hidden" type="submit">Submit</button>
	</form>
	
	{#if $message}<p>{$message}</p>{/if}

	{#if JSON.stringify(years) !== JSON.stringify([])}	
	<div class="flex w-full justify-evenly mb-8">
		<div class="w-1/3 max-w-[340px] min-h-80">
			<BarChart data={meanDataSet} />
		</div>
		<div class="w-1/3 max-w-80 min-h-80">
			<BarChart data={countDataSet} />
		</div>
	</div>
	{/if}
	
	{#if form?.records}
	<div class="w-3/4">
		<table class="table w-full mb-6">
			<thead>
				<th>Month</th>
				<th>Town</th>
				<th>Flat Type</th>
				<th>Model</th>
				<th>Lease Start Year</th>
				<th>Remaining Lease</th>
				<th>Resale Price</th>
			</thead>
			<tbody>
				{#each paginatedSource as record}
					<tr class="text-center">
						<td>{record.month}</td>
						<td>{record.town}</td>
						<td>{record.flatType}</td>
						<td>{record.model}</td>
						<td>{record.leaseStart}</td>
						<td>{record.remainingLease}</td>
						<td>{record.price}</td>
					</tr>
				{/each}
			</tbody>
		</table>
	
		<Paginator 
				bind:settings={paginationSettings}
				showFirstLastButtons={true}
				controlVariant="focus:bg-secondary-500"
			/>
	</div>
	{:else if $message && !form?.records}
		<h2>No Results Found...</h2>
	{/if}
</div>

