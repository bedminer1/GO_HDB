import { z } from 'zod'
import { message, superValidate } from 'sveltekit-superforms'
import { zod } from 'sveltekit-superforms/adapters'
import { fail } from '@sveltejs/kit'

const schema = z.object({
    month: z.string(),
    town: z.string().min(1),
    flatType: z.string().min(1),
    leaseStart: z.string(),
    remainingLease: z.string(),
    price: z.number(),
})

export const load = async () => {
    const form = await superValidate( zod(schema))

    return {
        form
    }

}

export const actions = {
    query: async ({ request }) => {
        const form = await superValidate(request, zod(schema))

        if (!form.valid) {
            return fail(400, { form });
        }

        const url = `http://127.0.0.1:8080/2017/records?town=${form.data.town}&flat_type=${form.data.flatType}&price=${form.data.price}`
        const response = await fetch(url)
        const data = await response.json() // data in the form of [HDBRecord[], Info]

        const info: Map<string, number[]> = data[1]
        const records: HDBRecord[] = data[0]

        let years: string[] = []
        let meanData: number[] = []
        let countData: number[] = []
        for (const [year, data] of Object.entries(info)) {
			years = [...years, year]
			meanData = [...meanData, data[0]]
			countData = [...countData, data[1]]
		}

        message(form, "Query Submitted")

        return {
            records,
            years,
            meanData,
            countData,
            form, 
        }

    }
}