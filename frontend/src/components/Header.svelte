<script lang="ts">
    import storyquiz from '$assets/images/storyquiz.svg';
    import storyline from '$assets/images/storyline.svg';
    import csod from '$assets/images/csod.svg';
    import * as Popover from '$lib/components/ui/popover';
    import * as Drawer from '$lib/components/ui/drawer';
    import { Button } from '$lib/components/ui/button';
    import pkg from '../../package.json';

    import { BuildTime } from '../../wailsjs/go/main/App.js';

    let showLicense = false;
    let showInfo = false;

    const copyright = import.meta.env.VITE_COPYRIGHT || 'Copyright 2024 Alok Mishra';
    const email = import.meta.env.VITE_EMAIL || 'alok@alokmishra.com';

    let buildDay = 'April 2024';

    (async () => {
        try {
            const buildTime = await BuildTime();
            console.log('buildTime:' + buildTime);

            const date = new Date(buildTime);
            buildDay = date.toLocaleDateString('en-US', {
                weekday: 'long',
                year: 'numeric',
                month: 'long',
                day: 'numeric',
            });
        } catch (error) {
            console.error(error);
        }
    })();
</script>

<header class="relative">
    <h1 class="text-4xl font-black mt-4 text-orange-600">StoryQuiz</h1>
    <h2 class="text-sm">(Quiz Exporter - v{pkg.version})</h2>
    <Popover.Root bind:open={showInfo}>
        <Popover.Trigger>
            <img class="w-40 xs:w-60 p-4" alt="StoryQuiz Logo" src={storyquiz} />
        </Popover.Trigger>
        <Popover.Content class="w-56 xs:w-80 !top-12">
            <div class="flex justify-center items-center flex-col xs:flex-row">
                <div>
                    <img class="w-20 p-4" alt="StoryQuiz Logo" src={storyquiz} />
                </div>
                <div>
                    <p class="text-lg font-bold text-orange-600">Story Quiz</p>
                    <h4 class="text-sm font-semibold text-cyan-200">
                        <a href={`mailto:${email}`}>{email}</a>
                    </h4>
                    <h6 class="text-xs text-muted-foreground">v{pkg.version}, {buildDay}</h6>
                    <button
                        class="text-sm text-cyan-600 hover:text-cyan-400"
                        on:click={() => {
                            showLicense = true;
                            showInfo = false;
                        }}
                    >
                        License
                    </button>
                </div>
            </div>
        </Popover.Content>
    </Popover.Root>
    <img class="w-40 p-4 absolute -bottom-0 -left-20 hidden xs:block" alt="Storyline Logo" src={storyline} />
    <img class="w-40 p-4 absolute -bottom-0 -right-20 hidden xs:block" alt="CSOD Logo" src={csod} />
</header>

<Drawer.Root bind:open={showLicense}>
    <Drawer.Content class="mx-10">
        <Drawer.Header>
            <Drawer.Title class="border-b border-solid mb-4 pb-2">MIT License</Drawer.Title>
            <Drawer.Description>
                <section class="flex flex-col gap-4">
                    <p>{copyright}</p>
                    <p>
                        Permission is hereby granted, free of charge, to any person obtaining a copy of this software
                        and associated documentation files (the “Software”), to deal in the Software without
                        restriction, including without limitation the rights to use, copy, modify, merge, publish,
                        distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the
                        Software is furnished to do so, subject to the following conditions:
                    </p>
                    <p>
                        The above copyright notice and this permission notice shall be included in all copies or
                        substantial portions of the Software.
                    </p>
                    <p>
                        THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING
                        BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
                        NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
                        DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
                        OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
                    </p>
                </section>
            </Drawer.Description>
        </Drawer.Header>
        <Drawer.Footer>
            <Drawer.Close><Button class="font-bold">Close</Button></Drawer.Close>
        </Drawer.Footer>
    </Drawer.Content>
</Drawer.Root>
