<script>
  import { onMount } from "svelte";
  import LabNav from "../../../../components/labNav.svelte";

  let doctors = []

   onMount(()=>{
    fetch("http://localhost:8000/doctor_list", {
      headers:{
        "content-type":"application/json",
        Authorization:`Bearer ${localStorage.getItem('hospital_token')}`
      }
    })
    .then(res=>res.json())
    .then(data=> {doctors=data; console.log(data)})
    .catch(err=>error=err)
   })

</script>

<LabNav />
<div class="px-10">
    <div class="flex flex-wrap space-x-10 gap-y-10 pt-10">
      {#each doctors as doctor(doctor.ID)}
        <a href={`doctor_list/${doctor.ID}`} class="bg-gray-100 h-64 w-48 rounded-2xl flex flex-col space-y-2 items-center  shadow-xl">
            <div class="flex h-52 bg-white w-full rounded-2xl">
                <img src={`http://localhost:8000/${doctor.tou}`}  alt="" srcset="" class="object-cover h-full w-full rounded-t-2xl">
            </div>     
            <div class="flex flex-col items-center">
              <p class="text-gray-700 font-semibold text-sm">Dr. {doctor.firstName} {doctor.lastName}</p>
              <p class="text-green-900 font-normal text-xs">{doctor.specialty}</p>
            </div>               
        </a>       
      {/each}                              
    </div>
</div>