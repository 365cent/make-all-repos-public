# Make all repos on github to public

The main.go is simple script written in golang to make visibility of all the repos on your Github account to public.

## Usage
Replace `<github username>` with your own GitHub username

Replace `<github personal access token>` with your own Personal access tokens (classic), which you can generate by visiting [https://github.com/settings/tokens](https://github.com/settings/tokens)

Run the script by executing the command `go run main.go` in your terminal

Once the script has finished executing, it should output a message indicating which repositories it successfully set to public.
