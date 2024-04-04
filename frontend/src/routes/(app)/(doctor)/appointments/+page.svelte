<script>
  import AppointmentsTab from "../../../../components/appointments_tab.svelte";
import DocNav from "../../../../components/docNav.svelte";
  import { onMount } from "svelte";

  let appointments = [];

  onMount(async()=>{
    await fetch("http://localhost:8000/doctor_appointments",{
        headers:{
            "content-type": "application/json",
            Authorization: `Bearer ${localStorage.getItem("hospital_token")}`
        }
    })
    .then(res => res.json())
    .then(data => {appointments=data; console.log(data)})
    .catch(err => console.log(err))
  })
</script>

<div class="bg-gray-200 h-screen grid grid-cols-12">
  <div class="col-span-1">
      <DocNav />
  </div>

  <div class="px-5 col-span-8 flex ">
    <div class="flex flex-wrap space-x-10 gap-y-10 pt-10 justify-center">
      {#each appointments as appointment(appointment.ID)}
        <a href={`appointments/appointment/${appointment.ID}`} class="h-36 w-36 flex flex-col space-y-2 items-center">
            <div class="flex h-52  w-full rounded-2xl">
                <img src={`http://localhost:8000/${appointment.patientProfile}`}  alt="" srcset="" class="object-cover h-full w-full rounded-full">
            </div>     
            <div class="flex flex-col items-center">
              <p class="text-gray-700 font-semibold text-sm">{appointment.patientFirstName} {appointment.patientLastName}</p>
              <p class="text-green-900 font-normal text-xs">{appointment.status}</p>
            </div>               
        </a> 

        {:else} 
        <div class="flex justify-center items-center">
          <p class="text-gray-700 font-medium flex items-center justify-center">No Appointment</p>         
        </div>     
        {/each}                              
    </div>
  </div>

  <div class="col-span-3">
    <AppointmentsTab />
  </div>
</div>