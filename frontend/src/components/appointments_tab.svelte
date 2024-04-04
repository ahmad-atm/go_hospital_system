<script>
    import { onMount } from "svelte";
    
    let appointments = []
        
    onMount(async()=>{
        await fetch("http://localhost:8000/limited_doctor_appointments",{
                headers:{
                    "content-type": "application/json",
                    Authorization: `Bearer ${localStorage.getItem("hospital_token")}`
                }
            })
            .then(res => res.json())
            .then(data => {appointments=data; console.log(appointments)})
            .catch(err => console.log(err))        
    })
</script>

<div class="col-span-2 scroll-my-3">
    <div class="bg-white h-screen p-8 flex flex-col space-y-8 items-center">
     <a href="/appointments" class="font-extrabold text-gray-700">Appointments</a>
        {#each appointments as appointment(appointment.ID)}
    <a href={`appointments/appointment/${appointment.ID}`} class="w-full flex">
        <div class=" w-full flex flex-col items-center justify-center border border-gray-200 py-8 rounded-lg">
            <p>{appointment.patientFirstName} {appointment.patientLastName}</p>
            <p class="">{appointment.status}</p>
        </div>                
    </a>                
    {:else}
        <p class="text-gray-700 font-medium flex items-center justify-center">No Appointment</p>                
    {/each}
    </div>
</div>