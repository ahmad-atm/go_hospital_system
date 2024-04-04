<script>
    import { onMount } from "svelte";

    let patientProfile = {}
    onMount(async()=>{
    await fetch("http://localhost:8000/patient_home",{
        headers:{
            "content-type": "application/json",
            Authorization: `Bearer ${localStorage.getItem("hospital_token")}`
        }
    })
    .then(res => res.json())
    .then(data => patientProfile=data)
    .catch(err => console.log(err))
    })
</script>


<div class="col-span-2">
    <div class="bg-white h-screen p-8 flex flex-col space-y-8 items-center shadow-gray-400">
        <p class="font-extrabold text-gray-700">My Profile</p>
        <div class="flex flex-col items-center justify-center rounded-full h-32 w-32 border b">
            <img src={`http://localhost:8000/${patientProfile.profileImage}`}  alt="" srcset="" class="object-cover h-full w-full rounded-full">
            <p class="font-bold text-md">{patientProfile.firstName}</p>                        
        </div>        
        <hr />   
        <!-- {#each patientProfile as patientProfile(patientProfile.ID)}
            <div class=" w-full flex flex-col items-center justify-center border border-gray-200 py-8 rounded-lg">
                    <p>{patientProfile.patientFirstName} {patientProfile.patientLastName}</p>
                    <p class="">{patientProfile.patient_id}</p>
            </div>                
        {/each} -->
    </div>                
</div>