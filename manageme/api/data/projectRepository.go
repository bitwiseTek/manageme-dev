package data

/**
 *
 * @author Sika Kay
 * @date 20/01/18
 *
 */
import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/bitwiseTek/manageme-dev/manageme/api/models"
)

//Project Struct for Mongo Persistence
type ProjectRepository struct {
	C *mgo.Collection
}

//AddTask persists Task associated with EmpID
func (r *ProjectRepository) AddTaskByEmpID(empID string) (task models.Task, err error) {
	objID := bson.NewObjectId()
	task.ID = objID
	task.empID = bson.ObjectIdHex(empID)
	task.Status = "Started"

	err = r.C.Insert(&task)
	return
}

//AddChildTask persists Task associated with EmpID, TaskID
func (r *ProjectRepository) AddChildTaskByEmpID(empID string, taskID string) (task models.Task, err error) {
	objID := bson.NewObjectId()
	task.ID = objID
	task.empID = bson.ObjectIdHex(empID)
	task.TaskDependent = bson.ObjectIdHex(taskID)

	err = r.C.Insert(&task)
	return
}

//GetTaskByID gets task associated with an ID
func (r *ProjectRepository) GetTaskByID(id string) (task models.Task, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&task)
	return
}

//GetTasks gets all tasks
func (r *ProjectRepository) GetTasks() []models.Task {
	var tasks []models.Task
	iter := r.C.Find(nil).Iter()
	result := models.Project{}
	for iter.Next(&result) {
		tasks = append(tasks, result)
	}
	return tasks
}

//GetTasksByOrgID gets tasks associated with an OrgID
func (r *ProjectRepository) GetTasksByOrgID(orgID string) []models.Task {
	var tasks []models.Task
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.Task{}
	for iter.Next(&result) {
		tasks = append(tasks, result)
	}
	return tasks
}

//GetTasksByEmpID gets tasks associated with an EmpID
func (r *ProjectRepository) GetTasksByEmpID(empID string) []models.Task {
	var tasks []models.Task
	empid := bson.ObjectIdHex(empID)
	iter := r.C.Find(bson.M{"empid": empid}).Iter()
	result := models.Task{}
	for iter.Next(&result) {
		tasks = append(tasks, result)
	}
	return tasks
}

//AddProjectTask persists project task associated with OrgID
func (r *ProjectRepository) AddProjectTaskByOrgID(orgID string) (task models.ProjectTask, err error) {
	objID := bson.NewObjectId()
	task.ID = objID
	task.OrgID = bson.ObjectIdHex(orgID)

	err = r.C.Insert(&task)
	return
}

//GetProjectTaskByID gets project task associated with an ID
func (r *ProjectRepository) GetProjectTaskByID(id string) (project models.ProjectTask, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&project)
	return
}

//GetProjectTasks gets all project tasks
func (r *ProjectRepository) GetProjectTasks() []models.ProjectTask {
	var projectTasks []models.ProjectTask
	iter := r.C.Find(nil).Iter()
	result := models.ProjectTask{}
	for iter.Next(&result) {
		tasks := result.Tasks
		for _, t = range tasks {
			projectTasks = append(t, result)
		}
	}
	return projectTasks
}

//GetProjectTasksByOrgID gets project tasks associated with an OrgID
func (r *ProjectRepository) GetProjectTasksByOrgID(orgID string) []models.ProjectTask {
	var projectTasks []models.ProjectTask
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.ProjectTask{}
	for iter.Next(&result) {
		tasks := result.Tasks
		for _, t = range tasks {
			projectTasks = append(t, result)
		}
	}
	return projectTasks
}

//AddProject persists project associated with MgrID
func (r *ProjectRepository) AddProjectMgrID(mgrID string, typeID string, projectTaskID string) (project models.Project, err error) {
	objID := bson.NewObjectId()
	project.ID = objID
	project.Manager = bson.ObjectIdHex(mgrID)
	project.ProjectTypeID = bson.ObjectIdHex(typeID)
	project.projectTaskID = bson.ObjectIdHex(projectTaskID)
	project.Status = "Not Started"

	err = r.C.Insert(&project)
	return
}

//GetProjectByID gets project associated with an ID
func (r *ProjectRepository) GetProjectByID(id string) (project models.Project, err error) {
	err = r.C.Find(bson.ObjectIdHex(id)).One(&project)
	return
}

//GetProjects gets all projects
func (r *ProjectRepository) GetProjects() []models.Project {
	var projects []models.Project
	iter := r.C.Find(nil).Iter()
	result := models.Project{}
	for iter.Next(&result) {
		projects = append(projects, result)
	}
	return projects
}

//GetProjectsByMgrID gets projects associated with an MgrID
func (r *ProjectRepository) GetProjectsByMgrID(mgrID string) []models.Project {
	var projects []models.Project
	mgrid := bson.ObjectIdHex(mgrID)
	iter := r.C.Find(bson.M{"mgrid": mgrid}).Iter()
	result := models.Project{}
	for iter.Next(&result) {
		projects = append(projects, result)
	}
	return projects
}

//GetProjectsByOrgID gets projects associated with an OrgID
func (r *ProjectRepository) GetProjectsByOrgID(orgID string) []models.Project {
	var projects []models.Project
	orgid := bson.ObjectIdHex(orgID)
	iter := r.C.Find(bson.M{"orgid": orgid}).Iter()
	result := models.Project{}
	for iter.Next(&result) {
		projects = append(projects, result)
	}
	return projects
}

//GetProjectsByTypeID gets projects associated with an TypeID
func (r *ProjectRepository) GetProjectsByTypeID(typeID string) []models.Project {
	var projects []models.Project
	typeid := bson.ObjectIdHex(typeID)
	iter := r.C.Find(bson.M{"typeid": typeid}).Iter()
	result := models.Project{}
	for iter.Next(&result) {
		projects = append(projects, result)
	}
	return projects
}