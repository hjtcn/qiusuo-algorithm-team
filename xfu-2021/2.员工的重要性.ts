/**
 * Definition for Employee.
 * class Employee {
 *     id: number
 *     importance: number
 *     subordinates: number
 *     constructor(id: number, importance: number, subordinates: number[]) {
 *         this.id = (id === undefined) ? 0 : id;
 *         this.importance = (importance === undefined) ? 0 : importance;
 *         this.subordinates = (subordinates === undefined) ? [] : subordinates;
 *     }
 * }
 */

function GetImportance(employees: Employee[], id: number): number {

    let employeeCount = (id: number): number => {
        for(let index = 0;index <= employees.length;index++){
            const { id:eId, importance, subordinates } = employees[index];
            if(eId === id) {
                return subordinates.reduce((prev, next) => prev + employeeCount(next), importance);
            }
        }
    }

    return employeeCount(id);
};

// 简单优化一下

function GetImportance1(employees: Employee[], id: number): number {

    for(let index = 0;index <= employees.length;index++){
        const { id:eId, importance, subordinates } = employees[index];
        
        if(eId === id) {
            return subordinates.reduce((prev, next) => prev + GetImportance(employees, next), importance);
        }
    }

};