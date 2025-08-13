## README.md

### Service for Generating and Downloading ZIP Archives

#### Description:
The service processes user requests containing external file links and generates a ZIP archive including these files. The system collects all specified files, compresses them into a single ZIP archive, and provides a download link back to the user.

---

### Features:

- Accepts user requests via HTTP API calls with an array of file links (default formats supported: JPEG, PDF).
- Downloads files from provided URLs.
- Assembles a ZIP archive consisting of the downloaded files.
- Outputs a direct download link for the resulting ZIP archive.
- Limits concurrent processing to up to 3 active tasks simultaneously.
- Restricts the maximum number of files per archive to 3.

---

### Technical Requirements:

- **Programming Language**: Go (Golang).
- **Web Framework**: Gin.

---

### Limitations:

- Maximum of **3 files** per single archive creation task.
- Processes only up to **3 parallel tasks** at any moment.

---

### Getting Started:

#### Download repository

```
git clone 
```

#### Setting Up the Configuration

Write the config variables in `.env` file.

Example `.env` content:

```
BIND_IP=127.0.0.1
LISTEN_PORT=8080
FILE_TYPES=image/jpeg,application/pdf
```

#### Launch The Service

```
go build -o link_service cmd/app/main.go
./link_service
```

#### Create New Task:

```
curl POST '127.0.0.1:8080/task'
```

#### Add Link To The Task:

```
curl '127.0.0.1:8080/task/1' \
--header 'Content-Type: application/json' \
--data '{
    "link":"https://example.com/link.jpg"
}'
```

#### Get Status For Current Tasks:

```
curl '127.0.0.1:8080/task/1/status'
```

the responce will contain a link to download the archive if there are 3 links

```
    "download_link": "download_link",
    "task": {
        "id": 1,
        "links": [
            "https://example.com/link.jpg",
            "https://example.com/link2.jpg",
            "https://example.com/link3.jpg"
        ],
        "status": "finished"
    }
```
