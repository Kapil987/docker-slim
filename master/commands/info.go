package commands

import (
	"fmt"

	"github.com/cloudimmunity/docker-slim/master/inspectors/image"
	"github.com/cloudimmunity/docker-slim/utils"

	log "github.com/Sirupsen/logrus"
	"github.com/cloudimmunity/go-dockerclientx"
	"github.com/dustin/go-humanize"
)

func OnInfo(imageRef string) {
	fmt.Println("docker-slim: [info] image=", imageRef)

	client, _ := docker.NewClientFromEnv()

	imageInspector, err := image.NewInspector(client, imageRef)
	utils.FailOn(err)

	log.Info("docker-slim: inspecting 'fat' image metadata...")
	err = imageInspector.Inspect()
	utils.FailOn(err)

	_, artifactLocation := utils.PrepareSlimDirs(imageInspector.ImageInfo.ID)
	imageInspector.ArtifactLocation = artifactLocation

	log.Infof("docker-slim: [%v] 'fat' image size => %v (%v)\n",
		imageInspector.ImageInfo.ID,
		imageInspector.ImageInfo.VirtualSize,
		humanize.Bytes(uint64(imageInspector.ImageInfo.VirtualSize)))

	log.Info("docker-slim: processing 'fat' image info...")
	err = imageInspector.ProcessCollectedData()
	utils.FailOn(err)

	fmt.Println("docker-slim: [info] done.")
}
