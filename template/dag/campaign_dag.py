import os
import sys
from datetime import datetime, timedelta

def generate_campaign_dag(dag_id, campaign_id, schedule, schedule_type):
    """Function to dynamically generate a campaign DAG based on the campaign inputs"""

    print(os.getcwd())

    default_args = {
        'owner': 'airflow',
        'depends_on_past': False,
        'retries': 1,
        'retry_delay': timedelta(minutes=1),
    }

    if schedule_type == 'one-time':
        start_date = datetime.strptime(schedule, "%Y-%m-%d %H:%M:%S")
        schedule_interval = None
    else:
        start_date = datetime.strptime(schedule, "%Y-%m-%d %H:%M:%S")
        schedule_interval = schedule

    # Defines DAG content
    dag_content = f"""
    from airflow import DAG
    from airflow.operators.dummy_operator import DummyOperator
    from datetime import datetime, timedelta

    default_args = {default_args}

    dag = DAG(
        dag_id='{dag_id}',
        default_args=default_args,
        schedule_interval='{schedule_interval}',
        start_date={start_date},
        catchup=False
    )
    
    campaignId = {campaign_id}

    start = DummyOperator(task_id="start", dag=dag)
    end = DummyOperator(task_id="end", dag=dag)

    start >> end
    """

    # Create DAG file in the dags directory
    print(os.getcwd())


    dags_directory = 'dags'
    if not os.path.exists(dags_directory):
        os.makedirs(dags_directory)

    dag_file_path = os.path.join(dags_directory, f"{dag_id}.py")

    with open(dag_file_path, 'w') as f:
        f.write(dag_content)

    print(f"DAG file {dag_file_path} created successfully.")

# Get campaign details passed from GitHub Actions inputs
if len(sys.argv) != 4:
    print("Usage: python generate_dag.py <campaign_id> <schedule> <schedule_type>")
    sys.exit(1)

dag_id = sys.argv[0]
campaign_id = sys.argv[1]
schedule = sys.argv[2]
schedule_type = sys.argv[3]

generate_campaign_dag(dag_id, campaign_id, schedule, schedule_type)