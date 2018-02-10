package controllers

/**
 *
 * @author Sika Kay
 * @date 07/02/18
 *
 */
 
import (
	"encoding/json"
	"net/http"

	httpcontext "github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"

	"github.com/bitwiseTek/manageme-dev/manageme/api/models"
	"github.com/bitwiseTek/manageme-dev/manageme/api/data"
	"github.com/bitwiseTek/manageme-dev/manageme/api/common"
)

//AddTask for /orgs/tasks/add api
func AddTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	empid = bson.ObjectIdHex(vars["empid"])
	var dataRes TaskResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid task data",
			500,
		)
		return
	}
	task := &dataRes.Data
	task.OrgID = orgid
	task.InitiatedBy = empid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("tasks")
	repo := &data.ProjectRepository{C: col}

	repo.AddTaskByEmpID(task)
	j, err := json.Marshal(TaskResource{Data: *task})
	if err != nil {
		common.DisplayAppError(
			w, 
			err, 
			"An unexpected error has occured", 
			500,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

//AddChildTask for /orgs/tasks/child/add api
func AddChildTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	empid = bson.ObjectIdHex(vars["empid"])
	taskid = bson.ObjectIdHex(vars["taskid"])
	var dataRes TaskResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid task data",
			500,
		)
		return
	}
	task := &dataRes.Data
	task.OrgID = orgid
	task.InitiatedBy = empid
	task.TaskDependent = taskid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("tasks")
	repo := &data.ProjectRepository{C: col}

	repo.AddChildTaskByEmpID(task)
	j, err := json.Marshal(TaskResource{Data: *task})
	if err != nil {
		common.DisplayAppError(
			w, 
			err, 
			"An unexpected error has occured", 
			500,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

//GetTasks for /orgs/tasks api
func GetTasks(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("tasks")
	repo := &data.ProjectRepository{C: col}
	accts := repo.GetAccounts()

	j, err := json.Marshal(TasksResource{Data: bills})
	if err != nil {
		common.DisplayAppError(
			w, 
			err, 
			"An unexpected error has occured", 
			500,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

//GetTasksByOrg for /orgs/tasks/org/{orgId} api
func GetTasksByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("tasks")
	repo := &data.ProjectRepository{C: col}
	tasks, err := repo.GetTasksByOrgID(orgid)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
		} else {
			common.DisplayAppError(
				w, 
				err, 
				"An unexpected error has occured",
				500,
			)
		}
		return
	}

	j, err := json.Marshal(TasksResource{Data: tasks})
	if err != nil {
		common.DisplayAppError(
			w, 
			err, 
			"An unexpected error has occured", 
			500,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

//GetTasksByTask for /orgs/tasks/task/{taskId} api
func GetTasksByTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskid = bson.ObjectIdHex(vars["taskid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("tasks")
	repo := &data.ProjectRepository{C: col}
	tasks, err := repo.GetTasksByTaskID(taskid)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
		} else {
			common.DisplayAppError(
				w, 
				err, 
				"An unexpected error has occured",
				500,
			)
		}
		return
	}

	j, err := json.Marshal(TasksResource{Data: tasks})
	if err != nil {
		common.DisplayAppError(
			w, 
			err, 
			"An unexpected error has occured", 
			500,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

//GetTasksByEmp for /orgs/tasks/employees/{empId} api
func GetTasksByEmp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empid = bson.ObjectIdHex(vars["empid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("tasks")
	repo := &data.ProjectRepository{C: col}
	tasks, err := repo.GetTasksByEmpID(empid)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
		} else {
			common.DisplayAppError(
				w, 
				err, 
				"An unexpected error has occured",
				500,
			)
		}
		return
	}

	j, err := json.Marshal(TasksResource{Data: tasks})
	if err != nil {
		common.DisplayAppError(
			w, 
			err, 
			"An unexpected error has occured", 
			500,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

//GetTask for /orgs/tasks/{id} api
func GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("tasks")
	repo := &data.ProjectRepository{C: col}
	task, err := repo.GetTaskByID(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
		} else {
			common.DisplayAppError(
				w,
				err,
				"An unexpected error has occured",
				500,
			)
		}
		return
	}
	j, err := json.Marshal(task)
	if err != nil {
		common.DisplayAppError(
			w, 
			err,
			"An unexpected error has occured",
			500,
		)
		return
	}
	w.Header().Set("Content-Application", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//EditTaskByEmp for /orgs/tasks/employees/edit/{id} api
func EditTaskByEmp(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource TaskResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid task data",
			500,
		)
		return
	}
	task := &dataResource.Data
	task.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("tasks")
	repo := &data.AccountRepository{C: col}
	
	err := repo.EditTaskByID(task); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occured",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

//Delete Task - TBA

//AddProjectTask for /orgs/projects/tasks/add api
func AddProjectTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	var dataRes ProjectTaskResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid task data",
			500,
		)
		return
	}
	projectTask := &dataRes.Data
	projectTask.OrgID = orgid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("projecttasks")
	repo := &data.ProjectRepository{C: col}

	repo.AddProjectTaskByOrgID(projectTask)
	j, err := json.Marshal(ProjectTaskResource{Data: *projectTask})
	if err != nil {
		common.DisplayAppError(
			w, 
			err, 
			"An unexpected error has occured", 
			500,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

//GetProjectTasks for /orgs/projects/tasks api
func GetProjectTasks(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("projecttasks")
	repo := &data.ProjectRepository{C: col}
	projectTasks := repo.GetProjectTasks()

	j, err := json.Marshal(ProjectTasksResource{Data: projectTasks})
	if err != nil {
		common.DisplayAppError(
			w, 
			err, 
			"An unexpected error has occured", 
			500,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

//GetProjectTasksByOrg for /orgs/projects/tasks/org/{orgId} api
func GetTasksByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("projecttasks")
	repo := &data.ProjectRepository{C: col}
	projectTasks, err := repo.GetProjectTasksByOrgID(orgid)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
		} else {
			common.DisplayAppError(
				w, 
				err, 
				"An unexpected error has occured",
				500,
			)
		}
		return
	}

	j, err := json.Marshal(ProjectTasksResource{Data: projectTasks})
	if err != nil {
		common.DisplayAppError(
			w, 
			err, 
			"An unexpected error has occured", 
			500,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

//GetProjectTask for /orgs/projects/tasks/{id} api
func GetProjectTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("projecttasks")
	repo := &data.ProjectRepository{C: col}
	projectTask, err := repo.GetProjectTaskByID(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
		} else {
			common.DisplayAppError(
				w,
				err,
				"An unexpected error has occured",
				500,
			)
		}
		return
	}
	j, err := json.Marshal(projectTask)
	if err != nil {
		common.DisplayAppError(
			w, 
			err,
			"An unexpected error has occured",
			500,
		)
		return
	}
	w.Header().Set("Content-Application", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//EditProjectTaskByEmp for /orgs/projects/tasks/edit/{id} api
func EditProjectTaskByEmp(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource ProjectTaskResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid task data",
			500,
		)
		return
	}
	projectTask := &dataResource.Data
	projectTask.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("projecttasks")
	repo := &data.AccountRepository{C: col}
	
	err := repo.EditProjectTaskByID(task); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occured",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

//Delete Project Task - TBA

//AddProjectByMgr for /orgs/projects/add api
func AddProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["ordid"])
	empid = bson.ObjectIdHex(vars["empid"])
	projecttaskid = bson.ObjectIdHex(vars["projecttaskid"])
	projecttypeid = bson.ObjectIdHex(vars["projecttypeid"])
	var dataRes ProjectResource

	err := json.NewDecoder(r.Body).Decode(&dataRes)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid project data",
			500,
		)
		return
	}
	project := &dataRes.Data
	project.OrgID = orgid
	project.Manager = empid
	project.ProjectTaskID = projecttaskid
	project.ProjectTypeID = projecttypeid
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("projects")
	repo := &data.ProjectRepository{C: col}

	repo.AddProjectByManager(project)
	j, err := json.Marshal(ProjectResource{Data: *project})
	if err != nil {
		common.DisplayAppError(
			w, 
			err, 
			"An unexpected error has occured", 
			500,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

//GetProjectsByOrg for /orgs/projects/org/{orgId} api
func GetProjectsByOrg(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgid = bson.ObjectIdHex(vars["orgid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("projects")
	repo := &data.ProjectRepository{C: col}
	projects, err := repo.GetProjectsByOrgID(orgid)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
		} else {
			common.DisplayAppError(
				w, 
				err, 
				"An unexpected error has occured",
				500,
			)
		}
		return
	}

	j, err := json.Marshal(ProjectsResource{Data: projects})
	if err != nil {
		common.DisplayAppError(
			w, 
			err, 
			"An unexpected error has occured", 
			500,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

//GetProjectsByMgr for /orgs/projects/employees/{empId} api
func GetProjectsByMgr(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empid = bson.ObjectIdHex(vars["empid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("projects")
	repo := &data.ProjectRepository{C: col}
	projects, err := repo.GetProjectsByMgrID(empid)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
		} else {
			common.DisplayAppError(
				w, 
				err, 
				"An unexpected error has occured",
				500,
			)
		}
		return
	}

	j, err := json.Marshal(ProjectsResource{Data: projects})
	if err != nil {
		common.DisplayAppError(
			w, 
			err, 
			"An unexpected error has occured", 
			500,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

//GetProjectsByType for /orgs/projects/types/{typeId} api
func GetProjectsByType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	typeid = bson.ObjectIdHex(vars["projecttypeid"])
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("projects")
	repo := &data.ProjectRepository{C: col}
	projects, err := repo.GetProjectsByMgrID(typeid)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
		} else {
			common.DisplayAppError(
				w, 
				err, 
				"An unexpected error has occured",
				500,
			)
		}
		return
	}

	j, err := json.Marshal(ProjectsResource{Data: projects})
	if err != nil {
		common.DisplayAppError(
			w, 
			err, 
			"An unexpected error has occured", 
			500,
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

//GetProject for /orgs/projects/{id} api
func GetProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id = vars["id"]
	context := NewContext()
	defer context.Close()
	col := context.dbCollection("projects")
	repo := &data.ProjectRepository{C: col}
	project, err := repo.GetProjectByID(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
		} else {
			common.DisplayAppError(
				w,
				err,
				"An unexpected error has occured",
				500,
			)
		}
		return
	}
	j, err := json.Marshal(project)
	if err != nil {
		common.DisplayAppError(
			w, 
			err,
			"An unexpected error has occured",
			500,
		)
		return
	}
	w.Header().Set("Content-Application", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//EditProjectByMgr for /orgs/projects/employees/edit/{id} api
func EditProjectByMgr(w http.ResponseWriter, *http.Request) {
	vars := mux.Vars(r)
	id = bson.ObjectIdHex(vars["id"])
	var dataResource ProjectResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid project data",
			500,
		)
		return
	}
	project := &dataResource.Data
	project.ID = id
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("projects")
	repo := &data.ProjectRepository{C: col}
	
	err := repo.EditProjectByID(project); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occured",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

//Delete Project - TBA