## Film Radar(FM) - Discover Your Personal Taste
Film Radar is our group project, an innovative Film Recommendation Software as a Service(SaaS) that
leverages advanced algorithms and user preferences to deliver personalized movie tastes.
In a time of plentiful entertainment options, our SaaS aims to simplify the movie selection
process, enhance user satisfaction, and drive engagement by offering tailored recommendations
that match individual tasted and preferences.

### Overview
![FM](./snapshots/FM.png)

### UI Service
The codes of UI service can be divided by two parts: React App Development and host server.

#### UI Development
To launch React App Development, make sure you have installed nodejs in your machine, then run:
```
cd esrs_ui
npm install
npm start
```
You would see a local server launched in local and listened to port 3000.

#### UI Deploy
First, build your UI by running:
```
cd esrs_ui
npm run build
```
Then there will be a folder called `build` generated. Move this folder to UI host server:
```
mv ./esrs_ui/build ./ui_host_server/static
```
After that, run:
```
cd ui_host_server
pip install -r requirements.txt
python app.py
```
The UI Service will be running on port 8081