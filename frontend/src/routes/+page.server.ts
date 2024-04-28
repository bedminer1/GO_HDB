import { z } from 'zod'
import { message, superValidate } from 'sveltekit-superforms'
import { zod } from 'sveltekit-superforms/adapters'
import { fail } from '@sveltejs/kit'

// schema for queries
const schema = z.object({
    month: z.string(),
    town: z.string(),
    flatType: z.string(),
    leaseStart: z.string(),
    price: z.number().default(20),
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

        // fetch from go
        const BASEURL = "http://127.0.0.1:8080/records"
        let queryParamsArr: string[] = []
        if (form.data.town !== "") {
            queryParamsArr = [...queryParamsArr, `town=${form.data.town.replace(' ', '+').toUpperCase()}` ]
        }
        if (form.data.flatType !== "") {
            queryParamsArr = [...queryParamsArr, `flatType=${form.data.flatType.replace(' ', '+').toUpperCase()}` ]
        }
        if (form.data.price !== 0) {
            queryParamsArr = [...queryParamsArr, `price=${form.data.price * 100000}` ]
        }
        const QUERYPARAMS = queryParamsArr.join("&")

        const response = await fetch(BASEURL + "?" + queryParamsArr.join("&"))
        const data = await response.json() // data in the form of [HDBRecord[], Stats]

        
        const info: Stats = data[1]
        const records: HDBRecord[] = data[0]
        
        // reordering info into arrays of years, meanData and countData
        let years: string[] = []
        let meanData: number[] = []
        let countData: number[] = []
        for (const [year, data] of Object.entries(info)) {
			years = [...years, year]
			meanData = [...meanData, data[0]]
			countData = [...countData, data[1]]
		}


        message(form, `Results for ${form.data.flatType} @${form.data.town}`)

        return {
            records,
            years,
            meanData,
            countData,
            form, 
        }

    }
}