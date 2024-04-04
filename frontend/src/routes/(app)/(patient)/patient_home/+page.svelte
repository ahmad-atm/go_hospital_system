<script>
    import { onMount } from "svelte";
    import PatNav from "../../../../components/patNav.svelte";
    import PatientTab from "../../../../components/patient_tab.svelte";
    import Icon from '@iconify/svelte';

    let patientProfile = {}
    let modelCount = {}

    onMount(async()=>{
        await fetch("http://localhost:8000/patient_home",{
            headers:{
                "content-type": "application/json",
                "Authorization": `Bearer ${localStorage.getItem("hospital_token")}`
            }
        })
        .then(res => res.json())
        .then(data => patientProfile=data)
        .catch(err => console.log(err))

        await fetch("http://localhost:8000/patient_model_count",{
            headers:{
                "content-type": "application/json",
                "Authorization": `Bearer ${localStorage.getItem("hospital_token")}`
            }
        })
        .then(res => res.json())
        .then(data => modelCount=data)
        .catch(err => console.log(err))
    })
</script>


<div class="bg-gray-200 h-screen grid grid-cols-12">
    <div class="col-span-1">
        <PatNav />
    </div>
    <div class="col-span-11">
        {#if patientProfile.firstName == ""}
            <div class="flex items-center justify-center min-h-screen">
                <a href="/patient_profile_register" class="bg-blue-800 hover:bg-blue-600 p-6 rounded-lg text-white font-medium">Create Profile</a> 
            </div>
        {:else}
            <div class="grid grid-cols-6">
                <div class="px-5 mt-7 flex flex-col space-y-6 col-span-4">
                    <div class="bg-blue-800  rounded-3xl py-16">
                        <div class="p-6 grid grid-cols-3">
                            <div class="col-span-2">
                                <div class="flex flex-col ml-6">
                                        <p class="text-white font-medium text-md">Welcome</p>
                                        <p class="text-white font-bold text-4xl">{patientProfile.firstName} {patientProfile.surname}</p>                        
                                </div>                        
                            </div>
        
                            <div class="flex items-center justify-center rounded-full h-32 w-32">
                                <img src={`http://localhost:8000/${patientProfile.profileImage}`}  alt="" srcset="" class="object-cover h-full w-full rounded-full">
                            </div>
                        </div>
                    </div>
                    <div class="flex space-x-4">
                        <a href={`/${patientProfile.ID}/appointments`} class="bg-teal-400 h-40 w-32 rounded-2xl p-5 flex flex-col space-y-3 items-center justify-center">                        
                            <Icon icon="teenyicons:appointments-solid" class="text-4xl text-teal-800"/>
                            <p class="text-teal-950 font-semibold">{modelCount.AppointmentCount}</p>
                            <p class="text-teal-950 font-semibold text-xs">Appointments</p>
                        </a>
                        <!-- <a href="/" class="bg-indigo-500 h-40 w-32 rounded-2xl p-5 flex flex-col items-center justify-center">                     -->
                            <!-- <Icon icon="dashicons:products"  class="text-4xl text-indigo-900"/> -->
                            <!-- <p class="text-indigo-900 font-semibold">{count.productCount}</p> -->
                        <!-- </a>      -->
                        <!-- <a href="/" class="bg-rose-400 h-40 w-32  rounded-2xl p-5 flex flex-col items-center justify-center">                     -->
                            <!-- <Icon icon="material-symbols:draft-orders-sharp"  class="text-4xl text-rose-800"/> -->
                            <!-- <p class="text-rose-900 font-semibold">{count.orderCount}</p> -->
                        <!-- </a>  -->
                        <!-- <a href="/" class="bg-gray-300 h-40 w-32  rounded-2xl p-5 flex flex-col items-center justify-center">                     -->
                            <!-- <Icon icon="typcn:plus"  class="text-4xl text-gray-700"/> -->
                            <!-- <p class="text-rose-900 font-semibold">02</p> -->
                        <!-- </a>                              -->
                    </div>   
                    <!-- <div> -->
                        <!-- <a href="/doctor_list" class="flex flex-col items-center">
                            <p class="text-2xl text-blue-600 font-extrabold">{modelCount.PatientTestCount}</p>
                            <p class="text-gray-700">Tests</p>
                        </a> -->
                        
                        <!-- <a href="/patient_list" class="flex flex-col items-center">
                            <p class="text-2xl text-blue-600 font-extrabold">#</p>
                            <p class="text-gray-700">#</p>
                        </a> -->
                    <!-- </div> -->
                </div>  
                <div class="col-span-2">
                    <PatientTab />
                  </div>

            </div>
        {/if}
    </div>


</div>