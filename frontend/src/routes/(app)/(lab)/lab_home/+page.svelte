<script>
    import { onMount } from "svelte";
    import LabNav from "../../../../components/labNav.svelte";

    let labProfile = {}
    let modelCount = {}
    let patientTestHistory = []

    onMount(async()=>{
        await fetch("http://localhost:8000/lab_home",{
            headers:{
                "content-type": "application/json",
                Authorization: `Bearer ${localStorage.getItem("hospital_token")}`
            }
        })
        .then(res => res.json())
        .then(data => labProfile=data)
        .catch(err => console.log(err))

        await fetch("http://localhost:8000/count",{
            headers:{
                "content-type": "application/json",
                Authorization: `Bearer ${localStorage.getItem("hospital_token")}`
            }
        })
        .then(res => res.json())
        .then(data => modelCount=data)
        .catch(err => console.log(err))

        await fetch("http://localhost:8000/patient_test_history",{
            headers:{
                "content-type": "application/json",
                Authorization: `Bearer ${localStorage.getItem("hospital_token")}`
            }
        })
        .then(res => res.json())
        .then(data => {patientTestHistory=data; console.log(patientTestHistory)})
        .catch(err => console.log(err))        
    })
</script>


<div class="bg-gray-100 h-screen">
    <LabNav />
    {#if labProfile.firstName == ""}
        <div class="flex items-center justify-center min-h-screen">
            <a href="/lab_profile_register" class="bg-rose-600 hover:bg-rose-500 p-6 rounded-lg text-white font-medium">Create Profile</a> 
        </div>
    {:else}
    <div class="">
        <div class="px-5 flex flex-col space-y-6 col-span-3">
            <div class="bg-rose-600 w-[60%] h-[30%] rounded-lg py-16 shadow-lg shadow-rose-200">
                <div class="p-6">
                    <div class="flex flex-col">
                        <p class="text-white font-medium text-md">Welcome</p>
                        <p class="text-white font-bold text-4xl">{labProfile.firstName} {labProfile.surname}</p>                        
                    </div>
                </div>
            </div>

            <div class="bg-white w-[60%] rounded-lg grid grid-cols-3 py-10 shadow-lg shadow-gray-300">
                <div class="flex flex-col items-center">
                    <p class="text-2xl text-rose-600 font-extrabold">{modelCount.LabAdminCount}</p>
                    <p class="text-gray-700">Lab Members</p>
                </div>

                <a href="/doctor_list" class="flex flex-col items-center">
                    <p class="text-2xl text-teal-600 font-extrabold">{modelCount.DoctorCount}</p>
                    <p class="text-gray-700">Doctors</p>
                </a>
                
                <a href="/patient_list" class="flex flex-col items-center">
                    <p class="text-2xl text-blue-600 font-extrabold">{modelCount.PatientCount}</p>
                    <p class="text-gray-700">Patients</p>
                </a>
            </div>
        </div>  
        
        <div class="fixed right-8 bg-white h-screen top-28 w-[30%] p-8 flex flex-col space-y-8 items-center rounded-2xl shadow-lg shadow-gray-400">
            <p class="font-extrabold text-gray-700 border-b-2 border-gray-600">Tests History</p>
            <!-- {#each patientTestHistory as history(history.ID)}
                <div class=" w-full flex flex-col items-center justify-center border border-gray-200 py-8 rounded-lg">
                        <p>{history.title}</p>
                        <p class="">{history.patientFirstName}</p>
                </div>   
                {:else}
                <p>No test History</p>             
            {/each} -->
        </div>
    </div>
        
    {/if}
</div>