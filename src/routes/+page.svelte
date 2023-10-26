<script lang="ts">
	interface Experience {
		company: string
		avatar: string
		position: string
		start: number
		end: number | null
		description: string
	}
	const experiences: Experience[] = [
		{
			company: "Stackd",
			avatar: "https://www.stackd.one/_next/image?url=https%3A%2F%2Fstackd-website.s3.us-east-2.amazonaws.com%2Fanalytic.png&w=3840&q=75",
			position: "Software Developer",
			start: 21,
			end: 22,
			description: "Built and maintained a CLI desktop application and Chrome extension using node.js, TS, Tailwind, and Sveltekit"

		}
	]

	const blogs = [
		{
			title: "Hobnob",
			description: "Find events in your city or favorite artists",
			date: "June 16, 2023",
			url: "https://github.com/moreSalt/hobnob-client"
		},
		{
			title: "Lever Scrape",
			description: "Scrape Lever and Greenhouse job boards for roles",
			date: "January 30, 2023",
			url: "https://github.com/moreSalt/lever-scrape"
		},
		{
			title: "HelloFresh to Instacart",
			description: "Easily find Hello Fresh recipe ingredients in Instacart ",
			date: "September 24, 2022",
			url: "https://github.com/moreSalt/hellofresh-instacart"
		}

	]

	const skills = [
		"JS/TS",
		"Python",
		"Go",
		"C/C++",
		"HTML/CSS",
		"Svelte",
		"Tailwind CSS",
		"AWS",
		"SQL",
		"MongoDB",
		"AWS",
		"Git"
	]
</script>


<div class="w-screen flex flex-row justify-center items-center bg-white">


	<div class="flex flex-col gap-8 w-screen max-w-7xl bg-white select-none p-8">


		<div class="rounded border-base-300 max-w-3xl h-fit mt-24">

			<div class="avatar">
				<div class="w-16 rounded-full">
				  <!-- svelte-ignore a11y-img-redundant-alt -->
				  <img src="https://avatars.githubusercontent.com/u/83256976?v=4" alt="profile picture"/>
				</div>
			  </div>
			<h1 class="text-[#27272a] font-bold text-5xl mb-6">Software Developer, Student, and Vitamin Collector</h1>
			<p class="text-base-300 leading-6">I'm Paul, A software developer and computer science student at University of Colorado Boulder</p>
			<div class="flex flex-row gap-2 items-center mt-6">
				<a href="https://github.com/moreSalt" target="_blank" rel="noreferrer" class="btn btn-ghost btn-square rounded-full">
					<svg xmlns="http://www.w3.org/2000/svg" height="1.25em" viewBox="0 0 496 512"><!--! Font Awesome Free 6.4.2 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. --><path d="M165.9 397.4c0 2-2.3 3.6-5.2 3.6-3.3.3-5.6-1.3-5.6-3.6 0-2 2.3-3.6 5.2-3.6 3-.3 5.6 1.3 5.6 3.6zm-31.1-4.5c-.7 2 1.3 4.3 4.3 4.9 2.6 1 5.6 0 6.2-2s-1.3-4.3-4.3-5.2c-2.6-.7-5.5.3-6.2 2.3zm44.2-1.7c-2.9.7-4.9 2.6-4.6 4.9.3 2 2.9 3.3 5.9 2.6 2.9-.7 4.9-2.6 4.6-4.6-.3-1.9-3-3.2-5.9-2.9zM244.8 8C106.1 8 0 113.3 0 252c0 110.9 69.8 205.8 169.5 239.2 12.8 2.3 17.3-5.6 17.3-12.1 0-6.2-.3-40.4-.3-61.4 0 0-70 15-84.7-29.8 0 0-11.4-29.1-27.8-36.6 0 0-22.9-15.7 1.6-15.4 0 0 24.9 2 38.6 25.8 21.9 38.6 58.6 27.5 72.9 20.9 2.3-16 8.8-27.1 16-33.7-55.9-6.2-112.3-14.3-112.3-110.5 0-27.5 7.6-41.3 23.6-58.9-2.6-6.5-11.1-33.3 2.6-67.9 20.9-6.5 69 27 69 27 20-5.6 41.5-8.5 62.8-8.5s42.8 2.9 62.8 8.5c0 0 48.1-33.6 69-27 13.7 34.7 5.2 61.4 2.6 67.9 16 17.7 25.8 31.5 25.8 58.9 0 96.5-58.9 104.2-114.8 110.5 9.2 7.9 17 22.9 17 46.4 0 33.7-.3 75.4-.3 83.6 0 6.5 4.6 14.4 17.3 12.1C428.2 457.8 496 362.9 496 252 496 113.3 383.5 8 244.8 8zM97.2 352.9c-1.3 1-1 3.3.7 5.2 1.6 1.6 3.9 2.3 5.2 1 1.3-1 1-3.3-.7-5.2-1.6-1.6-3.9-2.3-5.2-1zm-10.8-8.1c-.7 1.3.3 2.9 2.3 3.9 1.6 1 3.6.7 4.3-.7.7-1.3-.3-2.9-2.3-3.9-2-.6-3.6-.3-4.3.7zm32.4 35.6c-1.6 1.3-1 4.3 1.3 6.2 2.3 2.3 5.2 2.6 6.5 1 1.3-1.3.7-4.3-1.3-6.2-2.2-2.3-5.2-2.6-6.5-1zm-11.4-14.7c-1.6 1-1.6 3.6 0 5.9 1.6 2.3 4.3 3.3 5.6 2.3 1.6-1.3 1.6-3.9 0-6.2-1.4-2.3-4-3.3-5.6-2z"/></svg>
				</a>
				<a href="https://www.linkedin.com/in/paul-rodriguez-2b255b24a/" target="_blank" rel="noreferrer" class="btn btn-ghost btn-square rounded-full">
					<svg xmlns="http://www.w3.org/2000/svg" height="1.25em" viewBox="0 0 448 512"><!--! Font Awesome Free 6.4.2 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. --><path d="M416 32H31.9C14.3 32 0 46.5 0 64.3v383.4C0 465.5 14.3 480 31.9 480H416c17.6 0 32-14.5 32-32.3V64.3c0-17.8-14.4-32.3-32-32.3zM135.4 416H69V202.2h66.5V416zm-33.2-243c-21.3 0-38.5-17.3-38.5-38.5S80.9 96 102.2 96c21.2 0 38.5 17.3 38.5 38.5 0 21.3-17.2 38.5-38.5 38.5zm282.1 243h-66.4V312c0-24.8-.5-56.7-34.5-56.7-34.6 0-39.9 27-39.9 54.9V416h-66.4V202.2h63.7v29.2h.9c8.9-16.8 30.6-34.5 62.9-34.5 67.2 0 79.7 44.3 79.7 101.9V416z"/></svg>
				</a>
			</div>
		</div>
		
		<div class="w-full flex flex-row justify-between gap-4 mt-24" >
			<!-- BLOG -->
			<div class="w-full flex-grow gap-4">
				{#each blogs as blog}
					<a class="rounded w-full p-4 flex flex-col gap-2 hover:bg-zinc-100 hover:cursor-pointer" href={blog.url}>
						<p class="border-l-2 pl-4 text-sm text-zinc-400 border-zinc-400">{blog.date}</p>
						<h3 class="text-lg text-zinc-800 font-medium">{blog.title}</h3>
						<p class="text-zinc-600">{blog.description}</p>
						<a href={blog.url} class="link text-sm text-blue-500">More</a>

					</a>
				{/each}
			</div>

			<div class="w-full flex flex-col gap-4 max-w-md">
				<!-- Past work -->
				<div class="rounded border w-full p-4 gap-2 flex flex-col h-fit">
					<div class="w-full flex flex-row gap-4 items-center">
						<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
							<path stroke-linecap="round" stroke-linejoin="round" class="fill-zinc-200" d="M20.25 14.15v4.25c0 1.094-.787 2.036-1.872 2.18-2.087.277-4.216.42-6.378.42s-4.291-.143-6.378-.42c-1.085-.144-1.872-1.086-1.872-2.18v-4.25m16.5 0a2.18 2.18 0 00.75-1.661V8.706c0-1.081-.768-2.015-1.837-2.175a48.114 48.114 0 00-3.413-.387m4.5 8.006c-.194.165-.42.295-.673.38A23.978 23.978 0 0112 15.75c-2.648 0-5.195-.429-7.577-1.22a2.016 2.016 0 01-.673-.38m0 0A2.18 2.18 0 013 12.489V8.706c0-1.081.768-2.015 1.837-2.175a48.111 48.111 0 013.413-.387m7.5 0V5.25A2.25 2.25 0 0013.5 3h-3a2.25 2.25 0 00-2.25 2.25v.894m7.5 0a48.667 48.667 0 00-7.5 0M12 12.75h.008v.008H12v-.008z" />
						</svg>
						<p class="text-zinc-400">Work Experience</p>
					</div>

					{#each experiences as exp}
						<!-- svelte-ignore a11y-no-noninteractive-tabindex -->
						<div tabindex="0" class="collapse collapse-arrow  rounded hover:bg-zinc-100">
							<input type="checkbox" class="peer"/> 

							<div class="flex flex-row items-center gap-3 h-16 collapse-title peer-checked:bg-zinc-100">
					

								<div class="avatar">
									<div class="w-10 rounded-full ring ring-zinc-100 ring-offset-zinc-100 ring-offset-2">
									<img src={exp.avatar || "https://avatars.githubusercontent.com/u/83256976?v=4"} class="bg-zinc-100" alt="{exp.company} logo"/>
									</div>
								</div>

								<div class="flex flex-col w-full">
									<p class="font-semibold text-zinc-800 text-sm leading-6">{exp.company}</p>
									<div class="flex flex-row justify-between w-full">
										<p class="text-zinc-400 text-sm">{exp.position}</p>
										<p class="text-zinc-400 text-sm">{exp.start} - {exp.end || "present"}</p>
									</div>

								</div>
							</div>

							<div class="flex flex-row items-center gap-3 rounded  collapse-content peer-checked:bg-zinc-100">
								<p class="text-sm text-zinc-600">{exp.description}</p>
							</div>


						</div>


					{/each}
				</div>

				<div class="flex flex-col gap-4 rounded p-4 border">
					<div class="w-full flex flex-row gap-4 items-center">
						<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
							<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
								<path stroke-linecap="round" stroke-linejoin="round" class="fill-zinc-200" d="M11.42 15.17L17.25 21A2.652 2.652 0 0021 17.25l-5.877-5.877M11.42 15.17l2.496-3.03c.317-.384.74-.626 1.208-.766M11.42 15.17l-4.655 5.653a2.548 2.548 0 11-3.586-3.586l6.837-5.63m5.108-.233c.55-.164 1.163-.188 1.743-.14a4.5 4.5 0 004.486-6.336l-3.276 3.277a3.004 3.004 0 01-2.25-2.25l3.276-3.276a4.5 4.5 0 00-6.336 4.486c.091 1.076-.071 2.264-.904 2.95l-.102.085m-1.745 1.437L5.909 7.5H4.5L2.25 3.75l1.5-1.5L7.5 4.5v1.409l4.26 4.26m-1.745 1.437l1.745-1.437m6.615 8.206L15.75 15.75M4.867 19.125h.008v.008h-.008v-.008z" />
							  </svg>
							  
						</svg>
						<p class="text-zinc-400">Skills</p>
					</div>
					<div class="flex flex-row flex-wrap gap-2">
						{#each skills as skill}
						<div class="badge badge-ghost bg-zinc-100 border-zinc-100 text-zinc-400 text-xs">{skill}</div>
						{/each}
					</div>
		
				</div>
			</div>




		</div>
	</div>
</div>