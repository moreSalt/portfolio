const data = [
    {
        title: "QA Automation Intern",
        company: "Litify",
        dates: "Jan 2025 - Jun 2025",
        description:
            "Wrote test automation for a Salesforce platform using Cypress (JS) and then helped migrate over to Playwright (TS). Participated in the Agile software development cycle. ",
    },
    {
        title: "QA Intern",
        company: "CU Boulder",
        dates: "Jun 2024 - Dec 2024",
        description:
            "Performed manual tests of a variety of platforms. Wrote test automation in PyTest. Participated in the Agile software development cycle.",
    },
    {
        title: 'Co-Founder & Software Engineer',
        company: 'Stackd',
        dates: 'Apr 2021 - Oct 2022',
        description: 'I am very proud to have co-founded this. Here I wore many hats. From building out desktop applications in Node to writing our front-end in first vanilla HTML/CSS and then switching to Svelte. Learned how to receive feedback from customers, support customers, and dedicate yourself to a singular mission'
    }
];

import { ExternalLink } from "lucide-react";

export default function Page() {
    return (
        <div className="grid grid-cols-3 gap-4">
            {data.map((item, index) => (
                <div
                    className="aspect-square bg-card text-card-foreground flex flex-col gap-2 p-6 @container/card"
                    key={index}
                >
                    <h4 className="scroll-m-20 text-xl font-semibold tracking-tight">
                        {item.title}
                    </h4>
                    <p className="text-muted-foreground text-sm">
                        {item.company}, {item.dates}
                    </p>
                    <p className="leading-7">{item.description}</p>
                </div>
            ))}
        </div>
    );
}
