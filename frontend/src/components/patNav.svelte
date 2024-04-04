<script>
    import { onMount } from "svelte";
    import Icon from '@iconify/svelte';

    let patientProfile = {}

    onMount(()=>{
        fetch("http://localhost:8000/patient_home",{
            headers:{
                "content-type": "application/json",
                Authorization: `Bearer ${localStorage.getItem("hospital_token")}`
            }
        })
        .then(res => res.json())
        .then(data => {patientProfile=data})
        .catch(err => console.log(err))
    })
</script>

<nav class="flex flex-col h-screen justify-between items-center space-x-4 border-b-2 py-6 px-14 bg-blue-800">
    <div class="flex flex-col gap-y-10 items-center">
        <p class="text-3xl text-white font-extrabold flex items-center justify-center px-7">
            <Icon icon="fxemoji:hospital" />
        </p>

        <a href="/patient_home" class="text-white hover:underline">
            <Icon icon="majesticons:home" width=30/>  
        </a>        
        {#if patientProfile.firstName == ""}
            <!-- <a href="/patient_profile_register" class="text-white hover:underline">Create Profile</a> -->
        {:else if patientProfile.firstName != ""}
            <!-- <a href="/patient_tests" class="text-white hover:underline">My Tests</a>    -->
            <a href="/doctors" class="text-white hover:underline">
                <Icon icon="fa6-solid:user-doctor" width=30/>
            </a>  
            <a href={`/${patientProfile.ID}/appointments`} class="text-white hover:underline">
                <Icon icon="teenyicons:appointments-solid" width=30/>
            </a>   
        {/if}        
    </div>
    
    <div class="flex flex-col gap-y-7 items-center">
        {#if patientProfile.firstName != ""}
            <a href="/patient_profile" class="text-white hover:underline">
                <Icon icon="gg:profile" width=30/>
            </a>
        {/if}  
        <a href="/logout" class="text-white hover:underline">
            <Icon icon="heroicons-outline:logout" width=30 />
        </a>
    </div>    
</nav>