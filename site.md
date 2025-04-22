title: TFUtils-GO

---

#description

TFUtils-Go is a powerful Go rewrite of the TFUtils project, designed to help TFBern students manage their projects more efficiently. This new version brings improved performance, easier distribution, and maintains the user-friendly approach of the original - now powered by Go!

#main

# TFUtils-GO: Making Development Less Annoying

## The Idea: What It Is and Why I Built It

TFUtils-GO is the next evolution of my earlier project, TFUtils. The original version, written in Python, served its purpose but ran into some significant challenges, as detailed in the [TFUtils article](https://oseifert.vercel.app/projects/851103459). The two main hurdles were:

1.  **Difficult Distribution:** Sharing the Python tool with classmates required complex bundling into executables, which was often unreliable and platform-specific.
2.  **Network Speed:** Accessing project templates over the TFBern network via standard file sharing felt sluggish in the Python version.

While developing the original TFUtils, I started feeling constrained by its limitations and the growing complexity of managing the TUI and commands directly within the main application. This frustration pushed me to look for better solutions, which led me to discover and quickly grow fond of the Go programming language (Golang).

Go seemed ideal for tackling the distribution and potential performance issues head-on. Furthermore, I considered our specific context at TFBern: as Electronics Technician apprentices, many of us gain experience with C programming. In my opinion, the syntax and concepts in Go feel closer to C than Python does, potentially making it **easier for fellow students to understand and contribute** to the project compared to Python. Go also enforces a **strict, unified code style** through tools like `gofmt`, which is a huge advantage for collaboration and long-term maintenance, ensuring the codebase remains consistent even with multiple contributors.

A core goal driving this rewrite is **long-term maintainability**. I wanted to build a solid foundation that abstracts away the complexities of the TUI and command handling. This way, even after I leave TFBern, other students or teachers can hopefully **continue developing TFUtils-GO more easily** by focusing on creating specific command modules ("Charms") rather than wrestling with the core infrastructure.

As with the original, the primary audience for TFUtils-GO remains my fellow Electronics Technician apprentices at TFBern.

## The Journey: From Concept to Reality

Before diving into coding this time, I spent more time thinking about the architecture. I realized that simply rewriting TFUtils line-for-line in Go wouldn't address the underlying desire for a more modular and potentially reusable system, nor would it fully achieve the goal of easier future contributions. I wanted something that could potentially be adapted for other command-line automation tasks beyond just the TFBern context.

This led to the creation of **[Charmer](https://github.com/ImGajeed76/charmer)**, a separate Go library born out of the needs for TFUtils-GO. Charmer is still in its early stages, but it provides the core framework for building TUI-based command-line tools like TFUtils-GO. Its main job is similar to the discovery mechanism in the old TFUtils: it automatically finds command modules (which I call "Charms") located in a specific `Charms` folder, loads them, and presents them in the text-based interface. This modular approach is key to simplifying future development.

A significant amount of effort also went into optimizing file handling within Charmer. I developed a unified system for managing file paths, allowing it to treat SFTP paths (like `sftp://user@host/path/to/file`) and local paths (`C:\path\to\file` or `/path/to/file`) consistently. This makes operations like copying files between a remote SFTP server and the local machine (or even between two SFTP servers) much simpler to implement within a Charm.

Consequently, the TFUtils-GO project itself is much leaner this time. Instead of being packed with UI logic and command implementations, it primarily leverages the Charmer library for its core functionality. The specific TFBern-related tasks will be implemented as individual Charms, making the main project simpler and focusing contributions on self-contained modules.

## The Outcome: Where It Stands and What I Learned

TFUtils-GO is currently under active development and, while not yet matching the full feature set of the original Python version, the foundation is much stronger and designed for longevity. Development focus has understandably been on Charmer recently, as it provides the essential engine for TFUtils-GO and embodies the goal of creating an extensible core. However, the architecture is now in place, making it theoretically straightforward to reimplement all the previous TFUtils functions as Charms within this new Go-based structure.

This project has been a fantastic learning experience, significantly deepening my understanding of Go. While I'm still far from being an expert, I genuinely enjoy working with the language and its ecosystem, and I plan to continue using it for future projects. The focus on creating the reusable Charmer library has also taught me a lot about designing more abstract and maintainable software.

I'm particularly proud of the unified path handling system developed for Charmer and the overall modular design facilitated by the library. Although much of the visible work resides within Charmer for now, it directly serves the primary goals set for TFUtils-GO: improved performance, easier distribution, and crucially, a more sustainable and collaborative future for the tool at TFBern.
