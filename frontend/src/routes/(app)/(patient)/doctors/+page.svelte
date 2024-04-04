<script>
    import { onMount } from "svelte";
    import PatNav from "../../../../components/patNav.svelte";
    import PatientTab from "../../../../components/patient_tab.svelte";
  
    let doctors = []
  
     onMount(()=>{
      fetch("http://localhost:8000/doctor_list", {
        headers:{
          "content-type":"application/json",
          Authorization:`Bearer ${localStorage.getItem('hospital_token')}`
        }
      })
      .then(res=>res.json())
      .then(data=> doctors=data)
      .catch(err=>console.log(err))
     })
  
  </script>
  
<div class="bg-gray-200 h-screen grid grid-cols-12">
    <div class="col-span-1">
        <PatNav />
    </div>
    <div class="col-span-8 flex  justify-center">
      <div class="p-5">
        <div class="flex flex-wrap space-x-10 gap-y-10 pt-10">
            {#each doctors as doctor(doctor.ID)}
              <a href={`doctors/${doctor.ID}`} class="h-32 w-32 rounded-2xl flex flex-col space-y-2 items-center">
                <div class="flex h-52  w-full rounded-2xl">
                  <img src={`http://localhost:8000/${doctor.tou}`}  alt="" srcset="" class="object-cover h-full w-full rounded-full">
              </div>     
              <div class="flex flex-col items-center">
                <p class="text-gray-900 font-extrabold text-sm"> Dr. {doctor.firstName} {doctor.lastName}</p>
                <p class="text-blue-900 font-bold text-xs">{doctor.specialty}</p>
              </div>                             
              </a>       
            {/each}                              
        </div>      
      </div>      
    </div>
    <div class="col-span-3">
      <PatientTab />
    </div>
</div>
