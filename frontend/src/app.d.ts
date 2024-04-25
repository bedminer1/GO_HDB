// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
// and what to do when importing types
declare namespace App {
	// interface Locals {}
	// interface PageData {}
	// interface Error {}
	// interface Platform {}
}

interface HDBRecord {
	month: string
    town: string
    flatType: string
    leaseStart: string
    remainingLease: string
    price: number
}

type Info = Map<string, number[]>

interface dataSet {
    labels: string[];
    datasets: {
        label: string;
        data: number[];
        borderWidth: number;
        backgroundColor: string;
    }[];
}