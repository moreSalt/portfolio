<script lang="ts">

	const date1 = new Date("2020-08-01");
	const date2 = new Date();

    // One day in milliseconds
    const oneDay = 1000 * 60 * 60 * 24;

    // Calculating the time difference between two dates
    const diffInTime = date2.getTime() - date1.getTime();

    // Calculating the no. of days between two dates
    const diffInDays = Math.round(diffInTime / oneDay);

    const years = Math.round(diffInDays / 365)

	interface Work {
		title: string
		type: "work" |  "edu" | "project"
		at: string
		url: string
		startDate: string
		endDate: string
		description: string
	}

	interface Skill {
		title: string
		type: "lang" | "cloud" | "db" | "tool" | "styling"
	}

	const works: Work[] = [
		{
			title: "Software Developer",
			type: "work",
			at: "Redacted",
			url: "",
			startDate: "May 2021",
			endDate: "Present",
			description: `Built and maintained a Sveltekit + Tailwind CSS webapp. Worked on expanding features in the backend (Node.js). Currently working on moving everything to a serverless AWS environment and rewriting in Go`
		},
		{
			title: "Lever Scrape",
			type: "project",
			at: "Project",
			url: "https://github.com/moreSalt/lever-scrape",
			startDate: "Jan 2023",
			endDate: "Jan 2023",
			description: `A Lever and Greenhouse job board scraper built in Go to find positions based on user's keywords`
		},
		{
			title: "Portfolio",
			type: "project",
			at: "Project",
			url: "https://github.com/moreSalt/portfolio",
			startDate: "Nov 2022",
			endDate: "Present",
			description: `You are here right now. Built with Sveltekit and Tailwind, deployed on Vercel.`
		},
		{
			title: "Calenduh",
			type: "project",
			at: "Project",
			url: "https://github.com/moreSalt/calen-duh",
			startDate: "Oct 2022",
			endDate: "Oct 2022",
			description: `A simple web calendar I built to play with Sveltekit and Supabase.`
		},
		{
			title: "Hellofresh to Instacart",
			type: "project",
			at: "Project",
			url: "https://github.com/moreSalt/hellofresh-instacart",
			startDate: "Oct 2022",
			endDate: "Oct 2022",
			description: `A script to fetch all of the ingredients from a Hellofresh recipe and add them to Instacart`
		},
		{
			title: "Computer Science B.S.",
			type: "edu",
			at: "CU Boulder",
			url: "",
			startDate: "Aug 2022",
			endDate: "Present",
			description: `Getting educated one day at a time.`
		},
		{
			title: "Squarespace DL",
			type: "project",
			at: "Project",
			url: "",
			startDate: "Oct 2022",
			endDate: "Coming soon",
			description: `A package to scrape content from a Squarespace website, even if it's behind a paywall.`
		},
		{
			title: "+more",
			type: "project",
			at: "Project",
			url: "",
			startDate: "Aug 2020",
			endDate: "Present",
			description: `Wondering what else I have been working on, just shoot me an email.`
		},
	]

	const skills: Skill[] = [
		{
			title: "Javascript",
			type: "lang"
		},
		{
			title: "Typescript",
			type: "lang"
		},
		{
			title: "Node.js",
			type: "lang"
		},
		{
			title: "Svelte / Sveltekit",
			type: "lang",
		},
		{
			title: "Python",
			type: "lang"
		},
		{
			title: "Go",
			type: "lang"
		},
		{
			title: "C++",
			type: "lang"
		},
		{
			title: "HTML",
			type: "lang"
		},
		{
			title: "CSS",
			type: "styling"
		},
		{
			title: "Tailwind CSS",
			type: "styling"
		},
		{
			title: "Git",
			type: "tool"
		},
		{
			title: "AWS",
			type: "cloud"
		},
		{
			title: "Firebase",
			type: "cloud"
		},
		{
			title: "Supabase",
			type: "cloud"
		},
		{
			title: "MongoDB",
			type: "db"
		},
		{
			title: "Redis",
			type: "db"
		},
		{
			title: "PostgreSQL",
			type: "db"
		},
	]
</script>

<div class="navbar bg-base-100">
	<div class="navbar-start">
		<div class="dropdown">
			<!-- svelte-ignore a11y-no-noninteractive-tabindex -->
			<!-- svelte-ignore a11y-label-has-associated-control -->
			<label tabindex="0" class="btn btn-ghost btn-circle">
				<svg
					xmlns="http://www.w3.org/2000/svg"
					class="h-5 w-5"
					fill="none"
					viewBox="0 0 24 24"
					stroke="currentColor"
					><path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M4 6h16M4 12h16M4 18h7"
					/></svg
				>
			</label>
			<!-- svelte-ignore a11y-no-noninteractive-tabindex -->
			<ul
				tabindex="0"
				class="menu menu-compact dropdown-content border border-accent mt-3 p-2 shadow bg-base-100 rounded-lg w-52"
			>
				<li>
					<a href="https://github.com/moreSalt" target="_blank" rel="noreferrer">
						<svg
							xmlns="http://www.w3.org/2000/svg"
							fill="none"
							viewBox="0 0 24 24"
							stroke-width="1.5"
							stroke="currentColor"
							class="w-4 h-4"
						>
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								d="M13.19 8.688a4.5 4.5 0 011.242 7.244l-4.5 4.5a4.5 4.5 0 01-6.364-6.364l1.757-1.757m13.35-.622l1.757-1.757a4.5 4.5 0 00-6.364-6.364l-4.5 4.5a4.5 4.5 0 001.242 7.244"
							/>
						</svg>
						Github</a
					>
				</li>
				<li>
					<a href="https://gitlab.com/bannedSubnet" target="_blank" rel="noreferrer">
						<svg
							xmlns="http://www.w3.org/2000/svg"
							fill="none"
							viewBox="0 0 24 24"
							stroke-width="1.5"
							stroke="currentColor"
							class="w-4 h-4"
						>
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								d="M13.19 8.688a4.5 4.5 0 011.242 7.244l-4.5 4.5a4.5 4.5 0 01-6.364-6.364l1.757-1.757m13.35-.622l1.757-1.757a4.5 4.5 0 00-6.364-6.364l-4.5 4.5a4.5 4.5 0 001.242 7.244"
							/>
						</svg>
						Gitlab</a
					>
				</li>
			</ul>
		</div>
	</div>
	<div class="navbar-end">
		<a class="btn btn-ghost btn-circle" href="mailto:mildmoose@pm.me" target="_blank">
			<svg
				xmlns="http://www.w3.org/2000/svg"
				class="h-5 w-5"
				fill="none"
				viewBox="0 0 24 24"
				stroke="currentColor"
				><path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M21.75 6.75v10.5a2.25 2.25 0 01-2.25 2.25h-15a2.25 2.25 0 01-2.25-2.25V6.75m19.5 0A2.25 2.25 0 0019.5 4.5h-15a2.25 2.25 0 00-2.25 2.25m19.5 0v.243a2.25 2.25 0 01-1.07 1.916l-7.5 4.615a2.25 2.25 0 01-2.36 0L3.32 8.91a2.25 2.25 0 01-1.07-1.916V6.75"
				/></svg
			>
		</a>
	</div>
</div>

<main class="w-full flex flex-col items-center gap-8 mb-4 select-none">
	<p class="text-accent font-extrabold text-xl">Fullstack developer in Boulder.</p>

	<h1 class="text-5xl font-extrabold text-primary">Hi I'm Paul</h1>

	<!-- Short description -->
	<p class="w-4/5 text-center">
		I am a CS student at CU Boulder and a full-stack developer with {years}+ years of experience
	</p>

	<!-- SKILLS -->
	<div class="flex flex-wrap justify-center w-4/5 rounded gap-2">
		{#each skills as skill, i}
			<div class="badge badge-warning gap-2 items-center justify-center">
				<svg
					xmlns="http://www.w3.org/2000/svg"
					class="h-4 w-4"
					fill="none"
					viewBox="0 0 24 24"
					stroke="currentColor"
				>
					{#if skill.type === "lang"}
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M17.25 6.75L22.5 12l-5.25 5.25m-10.5 0L1.5 12l5.25-5.25m7.5-3l-4.5 16.5"
						/>
					{:else if skill.type === "db"}
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M20.25 6.375c0 2.278-3.694 4.125-8.25 4.125S3.75 8.653 3.75 6.375m16.5 0c0-2.278-3.694-4.125-8.25-4.125S3.75 4.097 3.75 6.375m16.5 0v11.25c0 2.278-3.694 4.125-8.25 4.125s-8.25-1.847-8.25-4.125V6.375m16.5 0v3.75m-16.5-3.75v3.75m16.5 0v3.75C20.25 16.153 16.556 18 12 18s-8.25-1.847-8.25-4.125v-3.75m16.5 0c0 2.278-3.694 4.125-8.25 4.125s-8.25-1.847-8.25-4.125"
						/>
					{:else if skill.type === "cloud"}
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M2.25 15a4.5 4.5 0 004.5 4.5H18a3.75 3.75 0 001.332-7.257 3 3 0 00-3.758-3.848 5.25 5.25 0 00-10.233 2.33A4.502 4.502 0 002.25 15z"
						/>
					{:else if skill.type === "styling"}
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M9.53 16.122a3 3 0 00-5.78 1.128 2.25 2.25 0 01-2.4 2.245 4.5 4.5 0 008.4-2.245c0-.399-.078-.78-.22-1.128zm0 0a15.998 15.998 0 003.388-1.62m-5.043-.025a15.994 15.994 0 011.622-3.395m3.42 3.42a15.995 15.995 0 004.764-4.648l3.876-5.814a1.151 1.151 0 00-1.597-1.597L14.146 6.32a15.996 15.996 0 00-4.649 4.763m3.42 3.42a6.776 6.776 0 00-3.42-3.42" 
						/>
					{:else if skill.type === "tool"}
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M21.75 6.75a4.5 4.5 0 01-4.884 4.484c-1.076-.091-2.264.071-2.95.904l-7.152 8.684a2.548 2.548 0 11-3.586-3.586l8.684-7.152c.833-.686.995-1.874.904-2.95a4.5 4.5 0 016.336-4.486l-3.276 3.276a3.004 3.004 0 002.25 2.25l3.276-3.276c.256.565.398 1.192.398 1.852z"
						/>
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M4.867 19.125h.008v.008h-.008v-.008z"
						/>
					{/if}

				</svg>
				{skill.title}
			</div>
		{/each}
	</div>

	<!-- WORKS -->
	<div class="grid grid-cols-1 gap-4 sm:grid-cols-2 w-4/5">

		{#each works as work, i}
			<a 
				class="border border-2 border-info shadow-xl rounded-lg p-4 w-full gap-2 flex flex-col"
				class:border-success="{i % 2 === 0}"
				class:border-warning="{i % 3 === 0}"
				class:border-error="{i % 4 === 0}"
				href={work.url}
				target="_blank"
				rel="noreferrer"
			>
				<p class="font-semibold text-sm flex gap-1 items-center">
					<svg
						xmlns="http://www.w3.org/2000/svg"
						fill="none"
						viewBox="0 0 24 24"
						stroke-width="1.5"
						stroke="currentColor"
						class="w-5 h-5"
					>

						{#if work.type === "work"}
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								d="M20.25 14.15v4.25c0 1.094-.787 2.036-1.872 2.18-2.087.277-4.216.42-6.378.42s-4.291-.143-6.378-.42c-1.085-.144-1.872-1.086-1.872-2.18v-4.25m16.5 0a2.18 2.18 0 00.75-1.661V8.706c0-1.081-.768-2.015-1.837-2.175a48.114 48.114 0 00-3.413-.387m4.5 8.006c-.194.165-.42.295-.673.38A23.978 23.978 0 0112 15.75c-2.648 0-5.195-.429-7.577-1.22a2.016 2.016 0 01-.673-.38m0 0A2.18 2.18 0 013 12.489V8.706c0-1.081.768-2.015 1.837-2.175a48.111 48.111 0 013.413-.387m7.5 0V5.25A2.25 2.25 0 0013.5 3h-3a2.25 2.25 0 00-2.25 2.25v.894m7.5 0a48.667 48.667 0 00-7.5 0M12 12.75h.008v.008H12v-.008z"
							/>
						{:else if work.type === "edu"}
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								d="M4.26 10.147a60.436 60.436 0 00-.491 6.347A48.627 48.627 0 0112 20.904a48.627 48.627 0 018.232-4.41 60.46 60.46 0 00-.491-6.347m-15.482 0a50.57 50.57 0 00-2.658-.813A59.905 59.905 0 0112 3.493a59.902 59.902 0 0110.399 5.84c-.896.248-1.783.52-2.658.814m-15.482 0A50.697 50.697 0 0112 13.489a50.702 50.702 0 017.74-3.342M6.75 15a.75.75 0 100-1.5.75.75 0 000 1.5zm0 0v-3.675A55.378 55.378 0 0112 8.443m-7.007 11.55A5.981 5.981 0 006.75 15.75v-1.5"
							/>

						{:else}
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								d="M12 18v-5.25m0 0a6.01 6.01 0 001.5-.189m-1.5.189a6.01 6.01 0 01-1.5-.189m3.75 7.478a12.06 12.06 0 01-4.5 0m3.75 2.383a14.406 14.406 0 01-3 0M14.25 18v-.192c0-.983.658-1.823 1.508-2.316a7.5 7.5 0 10-7.517 0c.85.493 1.509 1.333 1.509 2.316V18"
							/>
						{/if}
					</svg>
					{work.at}
				</p>
				<h2 
					class="card-title text-info"
					class:text-success="{i % 2 === 0}"
					class:text-warning="{i % 3 === 0}"
					class:text-error="{i % 4 === 0}"
				>{work.title}</h2>
				<p class="text-sm font-semibold">{work.startDate} - {work.endDate}</p>
				<p>
					{work.description}
				</p>
			</a>
		{/each}
	</div>
</main>

<footer class="footer footer-center p-4 bg-base-300 text-base-content">
    <div>
      <p>Made with the power of procrastination</p>
    </div>
</footer>

