import { z } from 'zod'
import { superValidate } from 'sveltekit-superforms'
import { zod } from 'sveltekit-superforms/adapters'
import { fail } from '@sveltejs/kit'

const schema = z.object({
    month: z.string(),
    town: z.string(),
    flatType: z.string(),
    leaseStart: z.string(),
    remainingLease: z.string(),
    price: z.number(),
})

export const load = async ({ fetch }) => {
    const form = await superValidate(zod(schema))

    try {
        const response = await fetch("http://localhost:8000/2017/records", { mode: 'no-cors' })
        const records = await response.json()

        return {
            form,
            records
        }
    } catch (err) {
        console.error(err)
    }

    return {
        form
    }
}

export const actions = {
    default: async ({ request }) => {
        const form = await superValidate(request, zod(schema))
        if (!form.valid) {
            return fail(400, { form });
        }

        return {
            form
        }
    }
}