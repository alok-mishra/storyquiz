<script lang="ts">
    // import logo from './assets/images/logo-universal.png';

    import storyquiz from './assets/images/storyquiz.svg';
    import storyline from './assets/images/storyline.svg';
    import csod from './assets/images/csod.svg';
    import excel from './assets/images/excel.svg';
    import word from './assets/images/word.svg';
    import csv from './assets/images/csv.svg';
    import './app.css';
    import logo from '../../build/appicon.png';
    import { Quiz } from '../wailsjs/go/main/App.js';

    import DragDrop from './components/DragDrop.svelte';

    import { Progress } from '$lib/components/ui/progress';
    import { Switch } from '$lib/components/ui/switch';
    import * as HoverCard from '$lib/components/ui/hover-card';

    let value = 66;

    let message: string = '👇 Please drop your file here 👇';
    let fileName: string;

    function onDrop(file: File | null, resultText: string): void {
        message = resultText;

        if (file) {
            fileName = file.name;
            const fileType = fileName.split('.').pop(); // file extension

            const reader = new FileReader();
            reader.onload = () => {
                const fileContents = reader.result as string;
                Quiz(btoa(fileContents), fileName, fileType).then((result) => {
                    message = result;
                });
            };
            reader.readAsBinaryString(file);
        }
    }
</script>

<main class="flex flex-col items-center overflow-auto">
    <!-- <img id="logo" class="w-40 p-4" alt="StoryQuiz Logo" src={logo} /> -->
    <h1 class="text-2xl font-black m-4">Quiz Exporter</h1>
    <section class="relative">
        <HoverCard.Root openDelay={2000}>
            <HoverCard.Trigger>
                <img class="w-60 p-4" alt="StoryQuiz Logo" src={storyquiz} />
            </HoverCard.Trigger>
            <HoverCard.Content class="w-60 !top-12">
                <div class="flex justify-around items-center">
                    <div>
                        <img class="w-20 p-4" alt="StoryQuiz Logo" src={storyquiz} />
                    </div>
                    <div>
                        <p class="text-lg">Quiz Exporter</p>
                        <h4 class="text-sm font-semibold">@alok.mishra</h4>
                        <span class="text-xs text-muted-foreground">v1.0, April 2024</span>
                    </div>
                </div>
            </HoverCard.Content>
        </HoverCard.Root>
        <img class="w-40 p-4 absolute -bottom-0 -left-20" alt="Storyline Logo" src={storyline} />
        <img class="w-40 p-4 absolute -bottom-0 -right-20" alt="CSOD Logo" src={csod} />
    </section>
    <Switch class="bg-orange-400" />
    <div id="message">{@html message}</div>
    <section class="dragdrop flex">
        <DragDrop dropColor={'bg-storyline'} fileTypes={['doc', 'docx', 'xls', 'xslx', 'xlsm']} {onDrop}>
            Drop a Word or Excel file here for Storyline Quiz
        </DragDrop>
        <DragDrop dropColor={'bg-csod'} fileTypes={['xls', 'xlsx', 'xlsm']} {onDrop} disabled>
            Drop a Excel file here for CSOD Quiz
        </DragDrop>
    </section>
    <footer class="w-full p-4">
        <Progress {value} max={100} class="max-w-60 mx-auto " />
    </footer>
</main>
