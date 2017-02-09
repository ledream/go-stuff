from invoke import task, run

@task
def build(ctx):
    ctx.run("git add --all")
    ctx.run("git commit -m clean")
    ctx.run("git push")
