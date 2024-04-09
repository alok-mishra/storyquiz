<script lang="ts">
    import storyquiz from '$assets/images/storyquiz.svg';
    import storyline from '$assets/images/storyline.svg';
    import csod from '$assets/images/csod.svg';
    import * as Popover from '$lib/components/ui/popover';
    import * as Drawer from '$lib/components/ui/drawer';
    import { Button } from '$lib/components/ui/button';

    let showLicense = false;
    let showInfo = false;
</script>

<section class="relative">
    <Popover.Root bind:open={showInfo}>
        <Popover.Trigger>
            <img class="w-40 xs:w-60 p-4" alt="StoryQuiz Logo" src={storyquiz} />
        </Popover.Trigger>
        <Popover.Content class="w-60 !top-12">
            <div class="flex justify-center items-center">
                <div>
                    <img class="w-20 p-4" alt="StoryQuiz Logo" src={storyquiz} />
                </div>
                <div>
                    <p class="text-lg font-bold text-orange-600">Story Quiz</p>
                    <h4 class="text-sm font-semibold text-cyan-200">@alok.mishra</h4>
                    <span class="text-xs text-muted-foreground">v1.0, April 2024</span>
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
</section>

<Drawer.Root bind:open={showLicense}>
    <Drawer.Content class="mx-10">
        <Drawer.Header>
            <Drawer.Title class="border-b border-solid mb-4 pb-2">MIT License</Drawer.Title>
            <Drawer.Description>
                <section class="flex flex-col gap-4">
                    <p>Copyright 2024 Alok Mishra</p>
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