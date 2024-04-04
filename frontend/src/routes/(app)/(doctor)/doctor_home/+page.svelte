<script>
    import { onMount } from "svelte";
    import DocNav from "../../../../components/docNav.svelte";
    import AppointmentsTab from "../../../../components/appointments_tab.svelte";

    let doctor_profile = {}
    let modelCount = []

    onMount(async()=>{
        await fetch("http://localhost:8000/doctor_home",{
            headers:{
                "content-type": "application/json",
                Authorization: `Bearer ${localStorage.getItem("hospital_token")}`
            }
        })
        .then(res => res.json())
        .then(data => doctor_profile=data)
        .catch(err => console.log(err))


        await fetch("http://localhost:8000/doctor_model_count",{
            headers:{
                "content-type": "application/json",
                Authorization: `Bearer ${localStorage.getItem("hospital_token")}`
            }
        })
        .then(res => res.json())
        .then(data => {modelCount=data; console.log(data)})
        .catch(err => console.log(err))
    })
</script>


<div class="bg-gray-200 h-screen grid grid-cols-12">
    <div class="col-span-1">
        <DocNav />
    </div>
    <div class="col-span-11">
        {#if doctor_profile.firstName == ""}
        <div class="flex items-center justify-center min-h-screen">
            <a href="/doctor_profile_register" class="bg-teal-600 hover:bg-teal-500 p-6 rounded-lg text-white font-medium">Create Profile</a> 
            </div>
        {:else}
        <div class="grid grid-cols-6">
            <div class="px-5 mt-7 flex flex-col space-y-6 col-span-4">
                <div class="bg-teal-800  rounded-3xl py-16">
                    <div class="p-6 grid grid-cols-3">
                        <div class="col-span-2">
                            <div class="flex flex-col ml-6">
                                <p class="text-white font-medium text-md">Welcome</p>
                                <p class="text-white font-bold text-4xl"> Dr. {doctor_profile.firstName} {doctor_profile.surname}</p>                                                                               
                            </div>
                        </div>
                        <div class="flex items-center justify-center rounded-full h-32 w-32">
                            <img src={`http://localhost:8000/${doctor_profile.tou}`}  alt="" srcset="" class="object-cover h-full w-full rounded-full">
                        </div>                        
                    </div>
                </div>

                <!-- <div class="bg-white w-[60%] rounded-lg grid grid-cols-3 py-10 shadow-lg shadow-gray-300">
                    <a href="/appointments" class="flex flex-col items-center">
                        <p class="text-2xl text-teal-600 font-extrabold">{modelCount.AppointmentCount}</p>
                        <p class="text-gray-700">Appointments</p>
                    </a> -->

                    <!-- <div class="flex flex-col items-center">
                        <p class="text-2xl text-teal-600 font-extrabold">#</p>
                        <p class="text-gray-700">Doctors</p>
                    </div> -->
                    
                    <!-- <a href="/doctor_tests" class="flex flex-col items-center">
                        <p class="text-2xl text-teal-600 font-extrabold">{modelCount.PatientCount}</p>
                        <p class="text-gray-700">Patients</p>
                    </a> -->
                <!-- </div> -->
                <div class="flex space-x-4 ">
                    <a href="/appointments" class="bg-teal-400 h-40 w-32 rounded-2xl p-5 flex flex-col items-center justify-center shadow-xl">                    
                        <!-- <Icon icon="ph:users-bold"  class="text-4xl text-teal-800"/> -->
                        <p class="text-2xl text-teal-600 font-extrabold">{modelCount.AppointmentCount}</p>
                        <p class="text-teal-900 font-semibold">Appointments</p>
                    </a>
                    <!-- <a href="/" class="bg-indigo-500 h-40 w-32 rounded-2xl p-5 flex flex-col items-center justify-center shadow-xl">                     -->
                        <!-- <Icon icon="dashicons:products"  class="text-4xl text-indigo-900"/> -->
                        <!-- <p class="text-indigo-900 font-semibold">{count.productCount}</p> -->
                    <!-- </a>      -->
                    <!-- <a href="/" class="bg-rose-400 h-40 w-32  rounded-2xl p-5 flex flex-col items-center justify-center shadow-xl">                     -->
                        <!-- <Icon icon="material-symbols:draft-orders-sharp"  class="text-4xl text-rose-800"/> -->
                        <!-- <p class="text-rose-900 font-semibold">{count.orderCount}</p> -->
                    <!-- </a>  -->
                    <!-- <a href="/" class="bg-gray-300 h-40 w-32  rounded-2xl p-5 flex flex-col items-center justify-center shadow-xl">                     -->
                        <!-- <Icon icon="typcn:plus"  class="text-4xl text-gray-700"/> -->
                        <!-- <p class="text-rose-900 font-semibold">02</p> -->
                    <!-- </a>                              -->
                </div>  
            </div>  
        <AppointmentsTab />

        </div>
        
        {/if}        
    </div>

</div>
