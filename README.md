# Bucket

Command-line tool to collect and build your own image bucket site.

## Usage

None of these actually exist yet, but this is the intention.

```
bucket add URL [filename]
```

Add the given image to your bucket. Optionally rename the added file.

```
bucket list
```

List the images in your bucket.

```
bucket build
```

Build the HTML for the bucket in a temporary directory.

```
bucket deploy
```

Build the HTML and deploy to your configured location.

```
bucket version
```

Display the version of the bucket command.

```
bucket init
```

Set up bucket for the first time. Creates the directory for which the images will go.

## TODOs

* Implement the commands documented above.
* Make this whole thing configurable.

