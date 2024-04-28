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
    model: string
    floorArea: number
    leaseStart: number
    remainingLease: string
    price: number
}

interface dataSet {
    labels: string[];
    datasets: {
        label: string;
        data: number[];
        borderWidth: number;
        backgroundColor: string;
    }[];
}

interface Stats {
    [key: string]: number[]
}