<script lang="ts">
    // @ts-nocheck
    import { Bar } from 'svelte-chartjs';
    import {Chart, Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale} from 'chart.js';
  
    Chart.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale);

    export let data: dataSet
  </script>
  
  <Bar 
  {data} 
  width={5}
  height={4}
  options={{
    animations: false,
    maintainAspectRatio: false,
    plugins: {
        legend: {
            labels: {
                color: '#DCC7EA'
            }
        }
    },
    scales: {
        y: {
            ticks: {
                color: '#DCC7EA'
            }
        },
        x: {
            ticks: {
                color: '#DCC7EA'
            }
        }
    },
}}/>