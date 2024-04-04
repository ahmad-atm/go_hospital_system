<script>
  import { onMount } from "svelte";
  import PatNav from "../../../../../../components/patNav.svelte";
  import {page} from "$app/stores";
  import PatientTab from "../../../../../../components/patient_tab.svelte";

  let appointmentId = $page.params.appointmentId;

  let appointment = {}

  onMount(async()=>{
        fetch(`http://localhost:8000/get_appointment/${appointmentId}`,{
        headers:{
            "Content-type":"application/json",
            Authorization: `Bearer ${localStorage.getItem('hospital_token')}`
        }
        })
        .then(res=>res.json())
        .then(data=>{appointment = data; console.log(data)})
        .catch(err=>console.log(err))
    })


</script>

<div class="bg-gray-200 h-screen grid grid-cols-12">
    <div class="col-span-1">
        <PatNav />
    </div>
    <div class="col-span-8 flex  justify-center">
        <div class="flex items-center justify-center space-x-5 gap-y-10 pt-4">
            <div class="h-64 w-48 rounded-2xl flex flex-col space-y-2 items-center">
                <div class="flex h-52 w-full rounded-2xl">
                    <img src={`http://localhost:8000/${appointment.doctorProfile}`}  alt="" srcset="" class="object-cover h-full w-full rounded-full">
                </div>     
                <div class="flex flex-col items-center">
                <p class="text-gray-700 font-semibold text-sm">{appointment.doctorFirstName}</p>
                <p class="text-gray-900 font-normal text-xs">{appointment.doctorLastName}</p>
                </div>               
            </div>
            <div>
                <p class="text-gray-700 font-semibold text-sm">First Name:  {appointment.doctorFirstName}</p>
                <p class="text-gray-700 font-semibold text-sm">Last Name:  {appointment.doctorLastName}</p>
                <p class="text-gray-700 font-semibold text-sm">Day:  {appointment.purpose}</p>
                <p class="text-blue-700 font-semibold text-sm">Status:  {appointment.status}</p>

            </div>        
        </div>
    <!-- <div class="flex items-center justify-center space-x-8 pt-7">      
        <button class="bg-rose-700 py-3 px-10 text-white rounded-lg hover:bg-rose-500">Decline</button>  
        <button class="bg-teal-700 py-3 px-10 text-white rounded-lg hover:bg-teal-500">Accept</button>        
    </div> -->
    </div>
    <div class="col-span-3">
        <PatientTab />
      </div>    
</div>