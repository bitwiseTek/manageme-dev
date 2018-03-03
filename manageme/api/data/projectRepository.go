package data

/**
 *
 * @author Sika Kay
 * @date 20/01/18
 *
 */
import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/bitwiseTek/manageme-dev/manageme/api/models"
)

//Project Struct for Mongo Persistence
type ProjectRepository struct {
	C *mgo.Collection
}

//AddTask persists Task associated with EmpID, OrgID
func (r *ProjectRepository) AddTaskByEmpID(orgID string, empID string) (task models.Task, err error) {
	objID := bson.NewObjectId()
	task.ID = objID
	task.OrgID = bson.ObjectIdHex(orgID)
	task.InitiatedBy = bson.ObjectIdHex(empID)
	task.Status = "Started"
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	err = r.C.Insert(&task)
	return
}

//AddChildTask persists Task associated with EmpID, TaskID, OrgID
func (r *ProjectRepository) AddChildTaskByEmpID(orgID string, empID string, taskID string) (task models.Task, err error) {
	objID := bson.NewObjectId()
	task.ID = objID
	task.OrgID = bson.ObjectIdHex(orgID)
	task.InitiatedBy = bson.ObjectIdHex(empID)
	task.TaskDependent = bson.ObjectIdHex(taskID)
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

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
	result := models.Task{}
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

//GetTasksByTaskID gets tasks associated with a TaskID
func (r *ProjectRepository) GetTasksByTaskID(taskID string) []models.Task {
	var tasks []models.Task
	taskid := bson.ObjectIdHex(taskID)
	iter := r.C.Find(bson.M{"taskid": taskid}).Iter()
	result := models.Task{}
	for iter.Next(&result) {
		tasks = append(tasks, result)
	}
	return tasks
}

//EditTaskByID edits task associated with an ID
func (r *ProjectRepository) EditTaskByID(task *models.Task) error {
	err := r.C.Update(bson.M{"_id": task.ID},
		bson.M{"$set": bson.M{
			"name":               task.Name,
			"subject":            task.Subject,
			"ismilestone":        task.IsMilestone,
			"priority":           task.Priority,
			"expectedstartdate":  task.ExpectedStartDate,
			"expectedenddate":    task.ExpectedEndDate,
			"actualstartdate":    task.ActualStartDate,
			"actualenddate":      task.ActualEndDate,
			"taskweight":         task.TaskWeight,
			"reviewdate":         task.ReviewDate,
			"note":               task.Note,
			"estimatedcosting":   task.EstimatedCosting,
			"totalcostingamount": task.TotalCostingAmount,
			"totalexpenseclaim":  task.TotalExpenseClaim,
			"totalbillingamount": task.TotalBillingAmount,
			"totalpurchasecost":  task.TotalPurchaseCost,
			"totalsalescost":     task.TotalSalesCost,
			"closingdate":        task.ClosingDate,
			"status":             task.Status,
			"updatedat":          time.Now(),
		}})
	return err
}

//DeleteTaskById deletes task out of the system by Id
func (r *ProjectRepository) DeleteTaskById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
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
		//tasks := result.Tasks
		/*for _, t := range tasks {
			t = append(t, result)
		}*/
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
		//tasks := result.Tasks
		/*for _, t := range tasks {
			t = append(t, result)
		}*/
	}
	return projectTasks
}

//EditProjectTaskByID edits project task associated with an ID
func (r *ProjectRepository) EditProjectTaskByID(task *models.ProjectTask) error {
	err := r.C.Update(bson.M{"_id": task.ID},
		bson.M{"$set": bson.M{
			"title":       task.Title,
			"description": task.Description,
			"tasks":       task.Tasks,
			"startdate":   task.StartDate,
			"enddate":     task.EndDate,
			"status":      task.Status,
		}})
	return err
}

//DeleteProjectTaskById deletes project task out of the system by Id
func (r *ProjectRepository) DeleteProjectTaskById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

//AddProject persists project associated with MgrID
func (r *ProjectRepository) AddProjectMgrID(orgID string, mgrID string, typeID string, projectTaskID string) (project models.Project, err error) {
	objID := bson.NewObjectId()
	project.ID = objID
	project.OrgID = bson.ObjectIdHex(orgID)
	project.Manager = bson.ObjectIdHex(mgrID)
	project.ProjectTypeID = bson.ObjectIdHex(typeID)
	project.ProjectTaskID = bson.ObjectIdHex(projectTaskID)
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

//EditProjectByID edits task associated with an ID
func (r *ProjectRepository) EditProjectByID(project *models.Project) error {
	err := r.C.Update(bson.M{"_id": project.ID},
		bson.M{"$set": bson.M{
			"name":               project.Name,
			"isactive":           project.IsActive,
			"priority":           project.Priority,
			"expectedstartdate":  project.ExpectedStartDate,
			"expectedenddate":    project.ExpectedEndDate,
			"actualstartdate":    project.ActualStartDate,
			"actualenddate":      project.ActualEndDate,
			"note":               project.Note,
			"estimatedcosting":   project.EstimatedCosting,
			"totalcostingamount": project.TotalCostingAmount,
			"totalexpenseclaim":  project.TotalExpenseClaim,
			"totalbillingamount": project.TotalBillingAmount,
			"totalpurchasecost":  project.TotalPurchaseCost,
			"totalsalescost":     project.TotalSalesCost,
			"status":             project.Status,
			"updatedat":          time.Now(),
		}})
	return err
}

//DeleteProjectById deletes project out of the system by Id
func (r *ProjectRepository) DeleteProjectById(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
