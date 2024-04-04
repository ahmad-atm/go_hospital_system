<script>
    import { onMount } from "svelte";
    import LabNav from "../../../../components/labNav.svelte";

    let patients = []

    onMount(()=>
     fetch("http://localhost:8000/patient_list",{
        headers:{
            "Content-type":"application/json",
            Authorization: `Bearer ${localStorage.getItem('hospital_token')}`
        }

     })
        .then(res=>res.json())
        .then(data=>{
            patients = data; 
            console.log(patients)
        })
        .catch(err=>console.log(err))
    )
</script>

<div>
    <LabNav />
    <div class="flex flex-col space-y-4">
        {#each patients as patient(patient.ID) }
        <div class="mx-24 py-10 p-5 border-2 border-rose-400 rounded-lg w-[50%] flex justify-between">
            <p  class="text-lg text-gray-800">{patient.firstName}  {patient.surname}</p>
            <a href={`/lab_create_test/${patient.ID}`} class="px-8 py-3 rounded-xl bg-rose-600 hover:bg-rose-400 text-white">Test</a>
        </div>
        {/each}        
    </div>

</div>